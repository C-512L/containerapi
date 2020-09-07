package containerapi

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	log "github.com/sirupsen/logrus"
)

type DockerProvider struct {
	client *client.Client
}

func (d DockerProvider) Create(ctx context.Context, cfg *ContainerConfig) (string, error) {
	dockerEnv := make([]string, 0, len(cfg.env))
	for k, v := range cfg.env {
		dockerEnv = append(dockerEnv, fmt.Sprintf("%s=%s", k, v))
	}

	config := &container.Config{
		Image: cfg.image,
		Env:   dockerEnv,
	}

	hostConfig := &container.HostConfig{
		AutoRemove:  true,
		NetworkMode: "host",
	}

	resp, err := d.client.ContainerCreate(ctx, config, hostConfig, nil, cfg.name)
	if err != nil {
		return "", fmt.Errorf("failed to create docker container: %w", err)
	}
	return resp.ID, nil
}

func (d DockerProvider) Start(ctx context.Context, containerID string) error {
	return d.client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
}

func (d DockerProvider) Stop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return d.client.ContainerStop(ctx, containerID, timeout)
}

func (d DockerProvider) Wait(ctx context.Context, containerID string) (<-chan int64, <-chan error) {
	msgChan := make(chan int64)
	errChan := make(chan error)

	go func() {
		defer close(msgChan)
		defer close(errChan)
		exit, err := d.client.ContainerWait(ctx, containerID)
		if err != nil {
			errChan <- err
		}
		msgChan <- exit
	}()

	return msgChan, errChan
}

func (d DockerProvider) Logs(ctx context.Context, containerID string) (io.ReadCloser, io.ReadCloser, error) {
	options := types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
	}
	combined, err := d.client.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get logs from docker: %w", err)
	}

	readOut, writeOut, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	readErr, writeErr, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	go func() {
		defer writeOut.Close()
		defer writeErr.Close()
		_, err := stdcopy.StdCopy(writeOut, writeErr, combined)
		if err != nil {
			log.WithField("err", err).Warn("failed to copy logs content to pipes")
		}
	}()

	return readOut, readErr, nil
}

func (d DockerProvider) Close() error {
	return d.client.Close()
}

func NewDockerProvider(ctx context.Context) (ContainerProvider, error) {
	apiclient, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %w", err)
	}

	_, err = apiclient.ServerVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection check failed: %w", err)
	}

	return &DockerProvider{
		client: apiclient,
	}, nil
}
