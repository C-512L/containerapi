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

// LibpodImageHistoryReader is a Reader for the LibpodImageHistory structure.
type LibpodImageHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodImageHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodImageHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodImageHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodImageHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodImageHistoryOK creates a LibpodImageHistoryOK with default headers values
func NewLibpodImageHistoryOK() *LibpodImageHistoryOK {
	return &LibpodImageHistoryOK{}
}

/* LibpodImageHistoryOK describes a response with status code 200, with default header values.

History response
*/
type LibpodImageHistoryOK struct {
	Payload *LibpodImageHistoryOKBody
}

func (o *LibpodImageHistoryOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/history][%d] libpodImageHistoryOK  %+v", 200, o.Payload)
}
func (o *LibpodImageHistoryOK) GetPayload() *LibpodImageHistoryOKBody {
	return o.Payload
}

func (o *LibpodImageHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageHistoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImageHistoryNotFound creates a LibpodImageHistoryNotFound with default headers values
func NewLibpodImageHistoryNotFound() *LibpodImageHistoryNotFound {
	return &LibpodImageHistoryNotFound{}
}

/* LibpodImageHistoryNotFound describes a response with status code 404, with default header values.

No such image
*/
type LibpodImageHistoryNotFound struct {
	Payload *LibpodImageHistoryNotFoundBody
}

func (o *LibpodImageHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/history][%d] libpodImageHistoryNotFound  %+v", 404, o.Payload)
}
func (o *LibpodImageHistoryNotFound) GetPayload() *LibpodImageHistoryNotFoundBody {
	return o.Payload
}

func (o *LibpodImageHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageHistoryNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImageHistoryInternalServerError creates a LibpodImageHistoryInternalServerError with default headers values
func NewLibpodImageHistoryInternalServerError() *LibpodImageHistoryInternalServerError {
	return &LibpodImageHistoryInternalServerError{}
}

/* LibpodImageHistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodImageHistoryInternalServerError struct {
	Payload *LibpodImageHistoryInternalServerErrorBody
}

func (o *LibpodImageHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/history][%d] libpodImageHistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodImageHistoryInternalServerError) GetPayload() *LibpodImageHistoryInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodImageHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImageHistoryInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodImageHistoryInternalServerErrorBody libpod image history internal server error body
swagger:model LibpodImageHistoryInternalServerErrorBody
*/
type LibpodImageHistoryInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod image history internal server error body
func (o *LibpodImageHistoryInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image history internal server error body based on context it is used
func (o *LibpodImageHistoryInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageHistoryInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageHistoryInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageHistoryInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImageHistoryNotFoundBody libpod image history not found body
swagger:model LibpodImageHistoryNotFoundBody
*/
type LibpodImageHistoryNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod image history not found body
func (o *LibpodImageHistoryNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image history not found body based on context it is used
func (o *LibpodImageHistoryNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageHistoryNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageHistoryNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageHistoryNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImageHistoryOKBody libpod image history o k body
swagger:model LibpodImageHistoryOKBody
*/
type LibpodImageHistoryOKBody struct {

	// comment
	Comment string `json:"Comment,omitempty"`

	// created
	Created int64 `json:"Created,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// ID
	ID string `json:"Id,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// tags
	Tags []string `json:"Tags"`
}

// Validate validates this libpod image history o k body
func (o *LibpodImageHistoryOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod image history o k body based on context it is used
func (o *LibpodImageHistoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImageHistoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImageHistoryOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodImageHistoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
