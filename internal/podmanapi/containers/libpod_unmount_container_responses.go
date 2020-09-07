// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodUnmountContainerReader is a Reader for the LibpodUnmountContainer structure.
type LibpodUnmountContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodUnmountContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLibpodUnmountContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodUnmountContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodUnmountContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodUnmountContainerNoContent creates a LibpodUnmountContainerNoContent with default headers values
func NewLibpodUnmountContainerNoContent() *LibpodUnmountContainerNoContent {
	return &LibpodUnmountContainerNoContent{}
}

/*LibpodUnmountContainerNoContent handles this case with default header values.

ok
*/
type LibpodUnmountContainerNoContent struct {
}

func (o *LibpodUnmountContainerNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] libpodUnmountContainerNoContent ", 204)
}

func (o *LibpodUnmountContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodUnmountContainerNotFound creates a LibpodUnmountContainerNotFound with default headers values
func NewLibpodUnmountContainerNotFound() *LibpodUnmountContainerNotFound {
	return &LibpodUnmountContainerNotFound{}
}

/*LibpodUnmountContainerNotFound handles this case with default header values.

No such container
*/
type LibpodUnmountContainerNotFound struct {
	Payload *LibpodUnmountContainerNotFoundBody
}

func (o *LibpodUnmountContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] libpodUnmountContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodUnmountContainerNotFound) GetPayload() *LibpodUnmountContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodUnmountContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodUnmountContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodUnmountContainerInternalServerError creates a LibpodUnmountContainerInternalServerError with default headers values
func NewLibpodUnmountContainerInternalServerError() *LibpodUnmountContainerInternalServerError {
	return &LibpodUnmountContainerInternalServerError{}
}

/*LibpodUnmountContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodUnmountContainerInternalServerError struct {
	Payload *LibpodUnmountContainerInternalServerErrorBody
}

func (o *LibpodUnmountContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] libpodUnmountContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodUnmountContainerInternalServerError) GetPayload() *LibpodUnmountContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodUnmountContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodUnmountContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodUnmountContainerInternalServerErrorBody libpod unmount container internal server error body
swagger:model LibpodUnmountContainerInternalServerErrorBody
*/
type LibpodUnmountContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod unmount container internal server error body
func (o *LibpodUnmountContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodUnmountContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodUnmountContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodUnmountContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodUnmountContainerNotFoundBody libpod unmount container not found body
swagger:model LibpodUnmountContainerNotFoundBody
*/
type LibpodUnmountContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod unmount container not found body
func (o *LibpodUnmountContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodUnmountContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodUnmountContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodUnmountContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
