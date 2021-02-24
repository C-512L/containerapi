// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodPingGetReader is a Reader for the LibpodPingGet structure.
type LibpodPingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodPingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodPingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodPingGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodPingGetOK creates a LibpodPingGetOK with default headers values
func NewLibpodPingGetOK() *LibpodPingGetOK {
	return &LibpodPingGetOK{}
}

/* LibpodPingGetOK describes a response with status code 200, with default header values.

Success
*/
type LibpodPingGetOK struct {

	/* Max compatibility API Version the server supports
	 */
	APIVersion string

	/* Default version of docker image builder
	 */
	BuildKitVersion string

	/* always no-cache
	 */
	CacheControl string

	/* If the server is running with experimental mode enabled, always true
	 */
	DockerExperimental bool

	/* Max Podman API Version the server supports.
	Available if service is backed by Podman, therefore may be used to
	determine if talking to Podman engine or another engine

	*/
	LibpodAPIVersion string

	/* Default version of libpod image builder.
	Available if service is backed by Podman, therefore may be used to
	determine if talking to Podman engine or another engine

	*/
	LibpodBuildhaVersion string

	/* always no-cache
	 */
	Pragma string

	Payload string
}

func (o *LibpodPingGetOK) Error() string {
	return fmt.Sprintf("[GET /libpod/_ping][%d] libpodPingGetOK  %+v", 200, o.Payload)
}
func (o *LibpodPingGetOK) GetPayload() string {
	return o.Payload
}

func (o *LibpodPingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header API-Version
	hdrAPIVersion := response.GetHeader("API-Version")

	if hdrAPIVersion != "" {
		o.APIVersion = hdrAPIVersion
	}

	// hydrates response header BuildKit-Version
	hdrBuildKitVersion := response.GetHeader("BuildKit-Version")

	if hdrBuildKitVersion != "" {
		o.BuildKitVersion = hdrBuildKitVersion
	}

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Docker-Experimental
	hdrDockerExperimental := response.GetHeader("Docker-Experimental")

	if hdrDockerExperimental != "" {
		valdockerExperimental, err := swag.ConvertBool(hdrDockerExperimental)
		if err != nil {
			return errors.InvalidType("Docker-Experimental", "header", "bool", hdrDockerExperimental)
		}
		o.DockerExperimental = valdockerExperimental
	}

	// hydrates response header Libpod-API-Version
	hdrLibpodAPIVersion := response.GetHeader("Libpod-API-Version")

	if hdrLibpodAPIVersion != "" {
		o.LibpodAPIVersion = hdrLibpodAPIVersion
	}

	// hydrates response header Libpod-Buildha-Version
	hdrLibpodBuildhaVersion := response.GetHeader("Libpod-Buildha-Version")

	if hdrLibpodBuildhaVersion != "" {
		o.LibpodBuildhaVersion = hdrLibpodBuildhaVersion
	}

	// hydrates response header Pragma
	hdrPragma := response.GetHeader("Pragma")

	if hdrPragma != "" {
		o.Pragma = hdrPragma
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodPingGetInternalServerError creates a LibpodPingGetInternalServerError with default headers values
func NewLibpodPingGetInternalServerError() *LibpodPingGetInternalServerError {
	return &LibpodPingGetInternalServerError{}
}

/* LibpodPingGetInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodPingGetInternalServerError struct {
	Payload *LibpodPingGetInternalServerErrorBody
}

func (o *LibpodPingGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/_ping][%d] libpodPingGetInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodPingGetInternalServerError) GetPayload() *LibpodPingGetInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodPingGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodPingGetInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodPingGetInternalServerErrorBody libpod ping get internal server error body
swagger:model LibpodPingGetInternalServerErrorBody
*/
type LibpodPingGetInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod ping get internal server error body
func (o *LibpodPingGetInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod ping get internal server error body based on context it is used
func (o *LibpodPingGetInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPingGetInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPingGetInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodPingGetInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
