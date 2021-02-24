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

// LibpodImagesImportReader is a Reader for the LibpodImagesImport structure.
type LibpodImagesImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodImagesImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodImagesImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLibpodImagesImportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodImagesImportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodImagesImportOK creates a LibpodImagesImportOK with default headers values
func NewLibpodImagesImportOK() *LibpodImagesImportOK {
	return &LibpodImagesImportOK{}
}

/* LibpodImagesImportOK describes a response with status code 200, with default header values.

Import response
*/
type LibpodImagesImportOK struct {
	Payload *LibpodImagesImportOKBody
}

func (o *LibpodImagesImportOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] libpodImagesImportOK  %+v", 200, o.Payload)
}
func (o *LibpodImagesImportOK) GetPayload() *LibpodImagesImportOKBody {
	return o.Payload
}

func (o *LibpodImagesImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImagesImportOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImagesImportBadRequest creates a LibpodImagesImportBadRequest with default headers values
func NewLibpodImagesImportBadRequest() *LibpodImagesImportBadRequest {
	return &LibpodImagesImportBadRequest{}
}

/* LibpodImagesImportBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type LibpodImagesImportBadRequest struct {
	Payload *LibpodImagesImportBadRequestBody
}

func (o *LibpodImagesImportBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] libpodImagesImportBadRequest  %+v", 400, o.Payload)
}
func (o *LibpodImagesImportBadRequest) GetPayload() *LibpodImagesImportBadRequestBody {
	return o.Payload
}

func (o *LibpodImagesImportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImagesImportBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImagesImportInternalServerError creates a LibpodImagesImportInternalServerError with default headers values
func NewLibpodImagesImportInternalServerError() *LibpodImagesImportInternalServerError {
	return &LibpodImagesImportInternalServerError{}
}

/* LibpodImagesImportInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodImagesImportInternalServerError struct {
	Payload *LibpodImagesImportInternalServerErrorBody
}

func (o *LibpodImagesImportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] libpodImagesImportInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodImagesImportInternalServerError) GetPayload() *LibpodImagesImportInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodImagesImportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImagesImportInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodImagesImportBadRequestBody libpod images import bad request body
swagger:model LibpodImagesImportBadRequestBody
*/
type LibpodImagesImportBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod images import bad request body
func (o *LibpodImagesImportBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod images import bad request body based on context it is used
func (o *LibpodImagesImportBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImagesImportBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImagesImportBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LibpodImagesImportBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImagesImportInternalServerErrorBody libpod images import internal server error body
swagger:model LibpodImagesImportInternalServerErrorBody
*/
type LibpodImagesImportInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod images import internal server error body
func (o *LibpodImagesImportInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod images import internal server error body based on context it is used
func (o *LibpodImagesImportInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImagesImportInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImagesImportInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodImagesImportInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImagesImportOKBody libpod images import o k body
swagger:model LibpodImagesImportOKBody
*/
type LibpodImagesImportOKBody struct {

	// Id
	ID string `json:"Id,omitempty"`
}

// Validate validates this libpod images import o k body
func (o *LibpodImagesImportOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod images import o k body based on context it is used
func (o *LibpodImagesImportOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImagesImportOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImagesImportOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodImagesImportOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}