// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodStopContainerReader is a Reader for the LibpodStopContainer structure.
type LibpodStopContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodStopContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLibpodStopContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLibpodStopContainerNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLibpodStopContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodStopContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodStopContainerNoContent creates a LibpodStopContainerNoContent with default headers values
func NewLibpodStopContainerNoContent() *LibpodStopContainerNoContent {
	return &LibpodStopContainerNoContent{}
}

/* LibpodStopContainerNoContent describes a response with status code 204, with default header values.

no error
*/
type LibpodStopContainerNoContent struct {
}

func (o *LibpodStopContainerNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] libpodStopContainerNoContent ", 204)
}

func (o *LibpodStopContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodStopContainerNotModified creates a LibpodStopContainerNotModified with default headers values
func NewLibpodStopContainerNotModified() *LibpodStopContainerNotModified {
	return &LibpodStopContainerNotModified{}
}

/* LibpodStopContainerNotModified describes a response with status code 304, with default header values.

Container already stopped
*/
type LibpodStopContainerNotModified struct {
	Payload *LibpodStopContainerNotModifiedBody
}

func (o *LibpodStopContainerNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] libpodStopContainerNotModified  %+v", 304, o.Payload)
}
func (o *LibpodStopContainerNotModified) GetPayload() *LibpodStopContainerNotModifiedBody {
	return o.Payload
}

func (o *LibpodStopContainerNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStopContainerNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodStopContainerNotFound creates a LibpodStopContainerNotFound with default headers values
func NewLibpodStopContainerNotFound() *LibpodStopContainerNotFound {
	return &LibpodStopContainerNotFound{}
}

/* LibpodStopContainerNotFound describes a response with status code 404, with default header values.

No such container
*/
type LibpodStopContainerNotFound struct {
	Payload *LibpodStopContainerNotFoundBody
}

func (o *LibpodStopContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] libpodStopContainerNotFound  %+v", 404, o.Payload)
}
func (o *LibpodStopContainerNotFound) GetPayload() *LibpodStopContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodStopContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStopContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodStopContainerInternalServerError creates a LibpodStopContainerInternalServerError with default headers values
func NewLibpodStopContainerInternalServerError() *LibpodStopContainerInternalServerError {
	return &LibpodStopContainerInternalServerError{}
}

/* LibpodStopContainerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodStopContainerInternalServerError struct {
	Payload *LibpodStopContainerInternalServerErrorBody
}

func (o *LibpodStopContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] libpodStopContainerInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodStopContainerInternalServerError) GetPayload() *LibpodStopContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodStopContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStopContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodStopContainerInternalServerErrorBody libpod stop container internal server error body
swagger:model LibpodStopContainerInternalServerErrorBody
*/
type LibpodStopContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stop container internal server error body
func (o *LibpodStopContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod stop container internal server error body based on context it is used
func (o *LibpodStopContainerInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStopContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStopContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodStopContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodStopContainerNotFoundBody libpod stop container not found body
swagger:model LibpodStopContainerNotFoundBody
*/
type LibpodStopContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stop container not found body
func (o *LibpodStopContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod stop container not found body based on context it is used
func (o *LibpodStopContainerNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStopContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStopContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodStopContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodStopContainerNotModifiedBody libpod stop container not modified body
swagger:model LibpodStopContainerNotModifiedBody
*/
type LibpodStopContainerNotModifiedBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stop container not modified body
func (o *LibpodStopContainerNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod stop container not modified body based on context it is used
func (o *LibpodStopContainerNotModifiedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStopContainerNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStopContainerNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res LibpodStopContainerNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
