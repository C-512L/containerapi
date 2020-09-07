// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodGetInfoReader is a Reader for the LibpodGetInfo structure.
type LibpodGetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodGetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodGetInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodGetInfoOK creates a LibpodGetInfoOK with default headers values
func NewLibpodGetInfoOK() *LibpodGetInfoOK {
	return &LibpodGetInfoOK{}
}

/*LibpodGetInfoOK handles this case with default header values.

Info
*/
type LibpodGetInfoOK struct {
	Payload *LibpodGetInfoOKBody
}

func (o *LibpodGetInfoOK) Error() string {
	return fmt.Sprintf("[GET /libpod/info][%d] libpodGetInfoOK  %+v", 200, o.Payload)
}

func (o *LibpodGetInfoOK) GetPayload() *LibpodGetInfoOKBody {
	return o.Payload
}

func (o *LibpodGetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodGetInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodGetInfoInternalServerError creates a LibpodGetInfoInternalServerError with default headers values
func NewLibpodGetInfoInternalServerError() *LibpodGetInfoInternalServerError {
	return &LibpodGetInfoInternalServerError{}
}

/*LibpodGetInfoInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodGetInfoInternalServerError struct {
	Payload *LibpodGetInfoInternalServerErrorBody
}

func (o *LibpodGetInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/info][%d] libpodGetInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodGetInfoInternalServerError) GetPayload() *LibpodGetInfoInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodGetInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodGetInfoInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodGetInfoInternalServerErrorBody libpod get info internal server error body
swagger:model LibpodGetInfoInternalServerErrorBody
*/
type LibpodGetInfoInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod get info internal server error body
func (o *LibpodGetInfoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBody Info is the overall struct that describes the host system
// running libpod/podman
swagger:model LibpodGetInfoOKBody
*/
type LibpodGetInfoOKBody struct {

	// registries
	Registries map[string]interface{} `json:"registries,omitempty"`

	// host
	Host *LibpodGetInfoOKBodyHost `json:"host,omitempty"`

	// store
	Store *LibpodGetInfoOKBodyStore `json:"store,omitempty"`

	// version
	Version *LibpodGetInfoOKBodyVersion `json:"version,omitempty"`
}

// Validate validates this libpod get info o k body
func (o *LibpodGetInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStore(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodGetInfoOKBody) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(o.Host) { // not required
		return nil
	}

	if o.Host != nil {
		if err := o.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBody) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(o.Store) { // not required
		return nil
	}

	if o.Store != nil {
		if err := o.Store.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "store")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBody) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(o.Version) { // not required
		return nil
	}

	if o.Version != nil {
		if err := o.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHost HostInfo describes the libpod host
swagger:model LibpodGetInfoOKBodyHost
*/
type LibpodGetInfoOKBodyHost struct {

	// arch
	Arch string `json:"arch,omitempty"`

	// buildah version
	BuildahVersion string `json:"buildahVersion,omitempty"`

	// c groups version
	CGroupsVersion string `json:"cgroupVersion,omitempty"`

	// c p us
	CPUs int64 `json:"cpus,omitempty"`

	// event logger
	EventLogger string `json:"eventLogger,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// kernel
	Kernel string `json:"kernel,omitempty"`

	// linkmode
	Linkmode string `json:"linkmode,omitempty"`

	// mem free
	MemFree int64 `json:"memFree,omitempty"`

	// mem total
	MemTotal int64 `json:"memTotal,omitempty"`

	// o s
	OS string `json:"os,omitempty"`

	// rootless
	Rootless bool `json:"rootless,omitempty"`

	// runtime info
	RuntimeInfo map[string]interface{} `json:"runtimeInfo,omitempty"`

	// swap free
	SwapFree int64 `json:"swapFree,omitempty"`

	// swap total
	SwapTotal int64 `json:"swapTotal,omitempty"`

	// uptime
	Uptime string `json:"uptime,omitempty"`

	// conmon
	Conmon *LibpodGetInfoOKBodyHostConmon `json:"conmon,omitempty"`

	// distribution
	Distribution *LibpodGetInfoOKBodyHostDistribution `json:"distribution,omitempty"`

	// id mappings
	IDMappings *LibpodGetInfoOKBodyHostIDMappings `json:"idMappings,omitempty"`

	// oci runtime
	OciRuntime *LibpodGetInfoOKBodyHostOciRuntime `json:"ociRuntime,omitempty"`

	// remote socket
	RemoteSocket *LibpodGetInfoOKBodyHostRemoteSocket `json:"remoteSocket,omitempty"`

	// slirp4netns
	Slirp4netns *LibpodGetInfoOKBodyHostSlirp4netns `json:"slirp4netns,omitempty"`
}

// Validate validates this libpod get info o k body host
func (o *LibpodGetInfoOKBodyHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConmon(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDistribution(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIDMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOciRuntime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteSocket(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlirp4netns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateConmon(formats strfmt.Registry) error {

	if swag.IsZero(o.Conmon) { // not required
		return nil
	}

	if o.Conmon != nil {
		if err := o.Conmon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "conmon")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateDistribution(formats strfmt.Registry) error {

	if swag.IsZero(o.Distribution) { // not required
		return nil
	}

	if o.Distribution != nil {
		if err := o.Distribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "distribution")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateIDMappings(formats strfmt.Registry) error {

	if swag.IsZero(o.IDMappings) { // not required
		return nil
	}

	if o.IDMappings != nil {
		if err := o.IDMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "idMappings")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateOciRuntime(formats strfmt.Registry) error {

	if swag.IsZero(o.OciRuntime) { // not required
		return nil
	}

	if o.OciRuntime != nil {
		if err := o.OciRuntime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "ociRuntime")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateRemoteSocket(formats strfmt.Registry) error {

	if swag.IsZero(o.RemoteSocket) { // not required
		return nil
	}

	if o.RemoteSocket != nil {
		if err := o.RemoteSocket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "remoteSocket")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyHost) validateSlirp4netns(formats strfmt.Registry) error {

	if swag.IsZero(o.Slirp4netns) { // not required
		return nil
	}

	if o.Slirp4netns != nil {
		if err := o.Slirp4netns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "slirp4netns")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHost) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHost) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostConmon ConmonInfo describes the conmon executable being used
swagger:model LibpodGetInfoOKBodyHostConmon
*/
type LibpodGetInfoOKBodyHostConmon struct {

	// package
	Package string `json:"package,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this libpod get info o k body host conmon
func (o *LibpodGetInfoOKBodyHostConmon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostConmon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostConmon) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostConmon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostDistribution DistributionInfo describes the host distribution
// for libpod
swagger:model LibpodGetInfoOKBodyHostDistribution
*/
type LibpodGetInfoOKBodyHostDistribution struct {

	// distribution
	Distribution string `json:"distribution,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this libpod get info o k body host distribution
func (o *LibpodGetInfoOKBodyHostDistribution) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostDistribution) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostDistribution) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostIDMappings IDMappings describe the GID and UID mappings
swagger:model LibpodGetInfoOKBodyHostIDMappings
*/
type LibpodGetInfoOKBodyHostIDMappings struct {

	// g ID map
	GIDMap []*LibpodGetInfoOKBodyHostIDMappingsGidmapItems0 `json:"gidmap"`

	// UID map
	UIDMap []*LibpodGetInfoOKBodyHostIDMappingsUidmapItems0 `json:"uidmap"`
}

// Validate validates this libpod get info o k body host ID mappings
func (o *LibpodGetInfoOKBodyHostIDMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGIDMap(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUIDMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodGetInfoOKBodyHostIDMappings) validateGIDMap(formats strfmt.Registry) error {

	if swag.IsZero(o.GIDMap) { // not required
		return nil
	}

	for i := 0; i < len(o.GIDMap); i++ {
		if swag.IsZero(o.GIDMap[i]) { // not required
			continue
		}

		if o.GIDMap[i] != nil {
			if err := o.GIDMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "idMappings" + "." + "gidmap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LibpodGetInfoOKBodyHostIDMappings) validateUIDMap(formats strfmt.Registry) error {

	if swag.IsZero(o.UIDMap) { // not required
		return nil
	}

	for i := 0; i < len(o.UIDMap); i++ {
		if swag.IsZero(o.UIDMap[i]) { // not required
			continue
		}

		if o.UIDMap[i] != nil {
			if err := o.UIDMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libpodGetInfoOK" + "." + "host" + "." + "idMappings" + "." + "uidmap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappings) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostIDMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostIDMappingsGidmapItems0 IDMap contains a single entry for user namespace range remapping. An array
// of IDMap entries represents the structure that will be provided to the Linux
// kernel for creating a user namespace.
swagger:model LibpodGetInfoOKBodyHostIDMappingsGidmapItems0
*/
type LibpodGetInfoOKBodyHostIDMappingsGidmapItems0 struct {

	// container ID
	ContainerID int64 `json:"container_id,omitempty"`

	// host ID
	HostID int64 `json:"host_id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this libpod get info o k body host ID mappings gidmap items0
func (o *LibpodGetInfoOKBodyHostIDMappingsGidmapItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappingsGidmapItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappingsGidmapItems0) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostIDMappingsGidmapItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostIDMappingsUidmapItems0 IDMap contains a single entry for user namespace range remapping. An array
// of IDMap entries represents the structure that will be provided to the Linux
// kernel for creating a user namespace.
swagger:model LibpodGetInfoOKBodyHostIDMappingsUidmapItems0
*/
type LibpodGetInfoOKBodyHostIDMappingsUidmapItems0 struct {

	// container ID
	ContainerID int64 `json:"container_id,omitempty"`

	// host ID
	HostID int64 `json:"host_id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this libpod get info o k body host ID mappings uidmap items0
func (o *LibpodGetInfoOKBodyHostIDMappingsUidmapItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappingsUidmapItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostIDMappingsUidmapItems0) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostIDMappingsUidmapItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostOciRuntime OCIRuntimeInfo describes the runtime (crun or runc) being
// used with podman
swagger:model LibpodGetInfoOKBodyHostOciRuntime
*/
type LibpodGetInfoOKBodyHostOciRuntime struct {

	// name
	Name string `json:"name,omitempty"`

	// package
	Package string `json:"package,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this libpod get info o k body host oci runtime
func (o *LibpodGetInfoOKBodyHostOciRuntime) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostOciRuntime) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostOciRuntime) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostOciRuntime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostRemoteSocket RemoteSocket describes information about the API socket
swagger:model LibpodGetInfoOKBodyHostRemoteSocket
*/
type LibpodGetInfoOKBodyHostRemoteSocket struct {

	// exists
	Exists bool `json:"exists,omitempty"`

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this libpod get info o k body host remote socket
func (o *LibpodGetInfoOKBodyHostRemoteSocket) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostRemoteSocket) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostRemoteSocket) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostRemoteSocket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyHostSlirp4netns SlirpInfo describes the slirp executable that
// is being being used.
swagger:model LibpodGetInfoOKBodyHostSlirp4netns
*/
type LibpodGetInfoOKBodyHostSlirp4netns struct {

	// executable
	Executable string `json:"executable,omitempty"`

	// package
	Package string `json:"package,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this libpod get info o k body host slirp4netns
func (o *LibpodGetInfoOKBodyHostSlirp4netns) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostSlirp4netns) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyHostSlirp4netns) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyHostSlirp4netns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyStore StoreInfo describes the container storage and its
// attributes
swagger:model LibpodGetInfoOKBodyStore
*/
type LibpodGetInfoOKBodyStore struct {

	// config file
	ConfigFile string `json:"configFile,omitempty"`

	// graph driver name
	GraphDriverName string `json:"graphDriverName,omitempty"`

	// graph options
	GraphOptions map[string]interface{} `json:"graphOptions,omitempty"`

	// graph root
	GraphRoot string `json:"graphRoot,omitempty"`

	// graph status
	GraphStatus map[string]string `json:"graphStatus,omitempty"`

	// run root
	RunRoot string `json:"runRoot,omitempty"`

	// volume path
	VolumePath string `json:"volumePath,omitempty"`

	// container store
	ContainerStore *LibpodGetInfoOKBodyStoreContainerStore `json:"containerStore,omitempty"`

	// image store
	ImageStore *LibpodGetInfoOKBodyStoreImageStore `json:"imageStore,omitempty"`
}

// Validate validates this libpod get info o k body store
func (o *LibpodGetInfoOKBodyStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainerStore(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodGetInfoOKBodyStore) validateContainerStore(formats strfmt.Registry) error {

	if swag.IsZero(o.ContainerStore) { // not required
		return nil
	}

	if o.ContainerStore != nil {
		if err := o.ContainerStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "store" + "." + "containerStore")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodGetInfoOKBodyStore) validateImageStore(formats strfmt.Registry) error {

	if swag.IsZero(o.ImageStore) { // not required
		return nil
	}

	if o.ImageStore != nil {
		if err := o.ImageStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodGetInfoOK" + "." + "store" + "." + "imageStore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStore) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyStoreContainerStore ContainerStore describes the quantity of containers in the
// store by status
swagger:model LibpodGetInfoOKBodyStoreContainerStore
*/
type LibpodGetInfoOKBodyStoreContainerStore struct {

	// number
	Number int64 `json:"number,omitempty"`

	// paused
	Paused int64 `json:"paused,omitempty"`

	// running
	Running int64 `json:"running,omitempty"`

	// stopped
	Stopped int64 `json:"stopped,omitempty"`
}

// Validate validates this libpod get info o k body store container store
func (o *LibpodGetInfoOKBodyStoreContainerStore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStoreContainerStore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStoreContainerStore) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyStoreContainerStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyStoreImageStore ImageStore describes the image store.  Right now only the number
// of images present
swagger:model LibpodGetInfoOKBodyStoreImageStore
*/
type LibpodGetInfoOKBodyStoreImageStore struct {

	// number
	Number int64 `json:"number,omitempty"`
}

// Validate validates this libpod get info o k body store image store
func (o *LibpodGetInfoOKBodyStoreImageStore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStoreImageStore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyStoreImageStore) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyStoreImageStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodGetInfoOKBodyVersion Version is an output struct for varlink
swagger:model LibpodGetInfoOKBodyVersion
*/
type LibpodGetInfoOKBodyVersion struct {

	// API version
	APIVersion int64 `json:"APIVersion,omitempty"`

	// built
	Built int64 `json:"Built,omitempty"`

	// built time
	BuiltTime string `json:"BuiltTime,omitempty"`

	// git commit
	GitCommit string `json:"GitCommit,omitempty"`

	// go version
	GoVersion string `json:"GoVersion,omitempty"`

	// os arch
	OsArch string `json:"OsArch,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this libpod get info o k body version
func (o *LibpodGetInfoOKBodyVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyVersion) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoOKBodyVersion) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoOKBodyVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
