package test

import (
	"bytes"
	"context"
	"github.com/LINBIT/containerapi"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func containerName(name string) string {
	return strings.ReplaceAll(name, "/", "-")
}

func TestRun(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/hello-world", nil)
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		returnCodeChan, errChan := provider.Wait(ctx, id)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

const (
	logTestScript = `
echo stdout1
echo stderr1 >&2
echo stdout2
echo stderr2 >&2
sleep 1 # ensures we don't miss the container lifecycle
`
	expectedLogStdout = "stdout1\nstdout2\n"
	expectedLogStderr = "stderr1\nstderr2\n"
)

func TestRunWithLogs(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("sh", "-ec", logTestScript))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte(expectedLogStdout), stdout, []byte(expectedLogStderr), stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestContainerLifecycle(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("false"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(1), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestCopyFrom(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		tmp, err := ioutil.TempDir("", "*")
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := os.RemoveAll(tmp)
			assert.NoError(t, err)
		})

		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("sleep", "1"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		err = provider.CopyFrom(ctx, id, "/bin/busybox", tmp)
		assert.NoError(t, err)

		assert.FileExists(t, tmp + "/busybox")
	})
}

func TestLifecycle(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("echo", "foobar"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		// Assert that the logs are accessible even after the container exits.
		time.Sleep(1 * time.Second)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte("foobar\n"), stdout, []byte{}, stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}

		// Assert we can still access logs after container completed
		stdout, stderr, err = provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte("foobar\n"), stdout, []byte{}, stderr)

		// Assert that calling stop on an already stopped container works without error
		timeout := 1 * time.Second
		err = provider.Stop(ctx, id, &timeout)
		assert.NoError(t, err)

		err = provider.Remove(ctx, id)
		assert.NoError(t, err)

		// Asserts that the container is unknown to the container runtime after Remove()
		time.Sleep(1 * time.Second)
		err = provider.Start(ctx, id)
		assert.Error(t, err)
	})
}

func TestRunWithCancel(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("tail", "-f", "/dev/null"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			timeout := 1 * time.Second
			err := provider.Stop(ctx, id, &timeout)
			assert.NoError(t, err)
			err = provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		waitCtx, cancel := context.WithCancel(ctx)
		returnCodeChan, errChan := provider.Wait(waitCtx, id)

		cancel()

		select {
		case r := <-returnCodeChan:
			assert.FailNow(t, "got return code for cancelled Wait(): %d", r)
		case err := <-errChan:
			assert.Error(t, err)
		}
	})
}

// Tries to run the test function on any provider, skipping those that are not available
// If no provider is available, the whole test will fail
func runOnAllProviders(t *testing.T, timeout time.Duration, lambda func(context.Context, containerapi.ContainerProvider, *testing.T)) {
	anySuccess := false
	for _, provider := range containerapi.Providers() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)

		success := t.Run(provider, func(t *testing.T) {
			provider, err := containerapi.NewProvider(ctx, provider)
			if err != nil {
				t.Skipf("failed to initialize provider: %v", err)
			}

			lambda(ctx, provider, t)
		})

		if success {
			anySuccess = true
		}
	}

	if !anySuccess {
		t.Fatal("expected (at least) one test to succeed")
	}
}

func assertIOEquals(t *testing.T, expectedStdout []byte, stdout io.ReadCloser, expectedStderr []byte, stderr io.ReadCloser) bool {
	stdoutBuff := bytes.Buffer{}
	stderrBuff := bytes.Buffer{}

	vg := sync.WaitGroup{}
	vg.Add(2)
	go func() {
		defer vg.Done()
		_, err := io.Copy(&stdoutBuff, stdout)
		assert.NoError(t, err)
	}()
	go func() {
		defer vg.Done()
		_, err := io.Copy(&stderrBuff, stderr)
		assert.NoError(t, err)
	}()

	vg.Wait()

	outEqual := assert.Equal(t, expectedStdout, stdoutBuff.Bytes())
	errEqual := assert.Equal(t, expectedStderr, stderrBuff.Bytes())
	return outEqual && errEqual
}