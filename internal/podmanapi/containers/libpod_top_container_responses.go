// Code generated by go-swagger; DO NOT EDIT.

package containers

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
	"github.com/go-openapi/validate"
)

// LibpodTopContainerReader is a Reader for the LibpodTopContainer structure.
type LibpodTopContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodTopContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodTopContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodTopContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodTopContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodTopContainerOK creates a LibpodTopContainerOK with default headers values
func NewLibpodTopContainerOK() *LibpodTopContainerOK {
	return &LibpodTopContainerOK{}
}

/* LibpodTopContainerOK describes a response with status code 200, with default header values.

List processes in container
*/
type LibpodTopContainerOK struct {
	Payload *LibpodTopContainerOKBody
}

func (o *LibpodTopContainerOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/top][%d] libpodTopContainerOK  %+v", 200, o.Payload)
}
func (o *LibpodTopContainerOK) GetPayload() *LibpodTopContainerOKBody {
	return o.Payload
}

func (o *LibpodTopContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodTopContainerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodTopContainerNotFound creates a LibpodTopContainerNotFound with default headers values
func NewLibpodTopContainerNotFound() *LibpodTopContainerNotFound {
	return &LibpodTopContainerNotFound{}
}

/* LibpodTopContainerNotFound describes a response with status code 404, with default header values.

No such container
*/
type LibpodTopContainerNotFound struct {
	Payload *LibpodTopContainerNotFoundBody
}

func (o *LibpodTopContainerNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/top][%d] libpodTopContainerNotFound  %+v", 404, o.Payload)
}
func (o *LibpodTopContainerNotFound) GetPayload() *LibpodTopContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodTopContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodTopContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodTopContainerInternalServerError creates a LibpodTopContainerInternalServerError with default headers values
func NewLibpodTopContainerInternalServerError() *LibpodTopContainerInternalServerError {
	return &LibpodTopContainerInternalServerError{}
}

/* LibpodTopContainerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodTopContainerInternalServerError struct {
	Payload *LibpodTopContainerInternalServerErrorBody
}

func (o *LibpodTopContainerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/top][%d] libpodTopContainerInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodTopContainerInternalServerError) GetPayload() *LibpodTopContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodTopContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodTopContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodTopContainerInternalServerErrorBody libpod top container internal server error body
swagger:model LibpodTopContainerInternalServerErrorBody
*/
type LibpodTopContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod top container internal server error body
func (o *LibpodTopContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod top container internal server error body based on context it is used
func (o *LibpodTopContainerInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodTopContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodTopContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodTopContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodTopContainerNotFoundBody libpod top container not found body
swagger:model LibpodTopContainerNotFoundBody
*/
type LibpodTopContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod top container not found body
func (o *LibpodTopContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod top container not found body based on context it is used
func (o *LibpodTopContainerNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodTopContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodTopContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodTopContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodTopContainerOKBody libpod top container o k body
swagger:model LibpodTopContainerOKBody
*/
type LibpodTopContainerOKBody struct {

	// Each process running in the container, where each is process is an array of values corresponding to the titles
	// Required: true
	Processes [][]string `json:"Processes"`

	// The ps column titles
	// Required: true
	Titles []string `json:"Titles"`
}

// Validate validates this libpod top container o k body
func (o *LibpodTopContainerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodTopContainerOKBody) validateProcesses(formats strfmt.Registry) error {

	if err := validate.Required("libpodTopContainerOK"+"."+"Processes", "body", o.Processes); err != nil {
		return err
	}

	return nil
}

func (o *LibpodTopContainerOKBody) validateTitles(formats strfmt.Registry) error {

	if err := validate.Required("libpodTopContainerOK"+"."+"Titles", "body", o.Titles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this libpod top container o k body based on context it is used
func (o *LibpodTopContainerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodTopContainerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodTopContainerOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodTopContainerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
