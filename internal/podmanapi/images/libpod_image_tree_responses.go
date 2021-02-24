// Code generated by go-swagger; DO NOT EDIT.

package images

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

// LibpodImageTreeReader is a Reader for the LibpodImageTree structure.
type LibpodImageTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodImageTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodImageTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodImageTreeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodImageTreeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodImageTreeOK creates a LibpodImageTreeOK with default headers values
func NewLibpodImageTreeOK() *LibpodImageTreeOK {
	return &LibpodImageTreeOK{}
}

/* LibpodImageTreeOK describes a response with status code 200, with default header values.

Image tree response
*/
type LibpodImageTreeOK struct {
	Payload *LibpodImageTreeOKBody
}

func (o *LibpodImageTreeOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/tree][%d] libpodImageTreeOK  %+v", 200, o.Payload)
}
func (o *LibpodImageTreeOK) GetPayload() *LibpodImageTreeOKBody {
	return o.Payload
}

func (o *LibpodImageTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageTreeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImageTreeNotFound creates a LibpodImageTreeNotFound with default headers values
func NewLibpodImageTreeNotFound() *LibpodImageTreeNotFound {
	return &LibpodImageTreeNotFound{}
}

/* LibpodImageTreeNotFound describes a response with status code 404, with default header values.

No such image
*/
type LibpodImageTreeNotFound struct {
	Payload *LibpodImageTreeNotFoundBody
}

func (o *LibpodImageTreeNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/tree][%d] libpodImageTreeNotFound  %+v", 404, o.Payload)
}
func (o *LibpodImageTreeNotFound) GetPayload() *LibpodImageTreeNotFoundBody {
	return o.Payload
}

func (o *LibpodImageTreeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageTreeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImageTreeInternalServerError creates a LibpodImageTreeInternalServerError with default headers values
func NewLibpodImageTreeInternalServerError() *LibpodImageTreeInternalServerError {
	return &LibpodImageTreeInternalServerError{}
}

/* LibpodImageTreeInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodImageTreeInternalServerError struct {
	Payload *LibpodImageTreeInternalServerErrorBody
}

func (o *LibpodImageTreeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/tree][%d] libpodImageTreeInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodImageTreeInternalServerError) GetPayload() *LibpodImageTreeInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodImageTreeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageTreeInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodImageTreeInternalServerErrorBody libpod image tree internal server error body
swagger:model LibpodImageTreeInternalServerErrorBody
*/
type LibpodImageTreeInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod image tree internal server error body
func (o *LibpodImageTreeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image tree internal server error body based on context it is used
func (o *LibpodImageTreeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageTreeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageTreeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageTreeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImageTreeNotFoundBody libpod image tree not found body
swagger:model LibpodImageTreeNotFoundBody
*/
type LibpodImageTreeNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod image tree not found body
func (o *LibpodImageTreeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image tree not found body based on context it is used
func (o *LibpodImageTreeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageTreeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageTreeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageTreeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImageTreeOKBody libpod image tree o k body
swagger:model LibpodImageTreeOKBody
*/
type LibpodImageTreeOKBody struct {

	// ID
	ID string `json:"id,omitempty"`

	// layers
	Layers []interface{} `json:"layers"`

	// size
	Size string `json:"size,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this libpod image tree o k body
func (o *LibpodImageTreeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image tree o k body based on context it is used
func (o *LibpodImageTreeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageTreeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageTreeOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageTreeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}