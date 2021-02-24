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

// LibpodExportImageReader is a Reader for the LibpodExportImage structure.
type LibpodExportImageReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *LibpodExportImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodExportImageOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodExportImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodExportImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodExportImageOK creates a LibpodExportImageOK with default headers values
func NewLibpodExportImageOK(writer io.Writer) *LibpodExportImageOK {
	return &LibpodExportImageOK{

		Payload: writer,
	}
}

/* LibpodExportImageOK describes a response with status code 200, with default header values.

no error
*/
type LibpodExportImageOK struct {
	Payload io.Writer
}

func (o *LibpodExportImageOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/get][%d] libpodExportImageOK  %+v", 200, o.Payload)
}
func (o *LibpodExportImageOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *LibpodExportImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodExportImageNotFound creates a LibpodExportImageNotFound with default headers values
func NewLibpodExportImageNotFound() *LibpodExportImageNotFound {
	return &LibpodExportImageNotFound{}
}

/* LibpodExportImageNotFound describes a response with status code 404, with default header values.

No such image
*/
type LibpodExportImageNotFound struct {
	Payload *LibpodExportImageNotFoundBody
}

func (o *LibpodExportImageNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/get][%d] libpodExportImageNotFound  %+v", 404, o.Payload)
}
func (o *LibpodExportImageNotFound) GetPayload() *LibpodExportImageNotFoundBody {
	return o.Payload
}

func (o *LibpodExportImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodExportImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodExportImageInternalServerError creates a LibpodExportImageInternalServerError with default headers values
func NewLibpodExportImageInternalServerError() *LibpodExportImageInternalServerError {
	return &LibpodExportImageInternalServerError{}
}

/* LibpodExportImageInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodExportImageInternalServerError struct {
	Payload *LibpodExportImageInternalServerErrorBody
}

func (o *LibpodExportImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/get][%d] libpodExportImageInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodExportImageInternalServerError) GetPayload() *LibpodExportImageInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodExportImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodExportImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodExportImageInternalServerErrorBody libpod export image internal server error body
swagger:model LibpodExportImageInternalServerErrorBody
*/
type LibpodExportImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod export image internal server error body
func (o *LibpodExportImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod export image internal server error body based on context it is used
func (o *LibpodExportImageInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodExportImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodExportImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodExportImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodExportImageNotFoundBody libpod export image not found body
swagger:model LibpodExportImageNotFoundBody
*/
type LibpodExportImageNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod export image not found body
func (o *LibpodExportImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod export image not found body based on context it is used
func (o *LibpodExportImageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodExportImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodExportImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodExportImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
