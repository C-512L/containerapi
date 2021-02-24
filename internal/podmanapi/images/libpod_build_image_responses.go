// Code generated by go-swagger; DO NOT EDIT.

package images

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

// LibpodBuildImageReader is a Reader for the LibpodBuildImage structure.
type LibpodBuildImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodBuildImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodBuildImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLibpodBuildImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodBuildImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodBuildImageOK creates a LibpodBuildImageOK with default headers values
func NewLibpodBuildImageOK() *LibpodBuildImageOK {
	return &LibpodBuildImageOK{}
}

/* LibpodBuildImageOK describes a response with status code 200, with default header values.

OK (As of version 1.xx)
*/
type LibpodBuildImageOK struct {
	Payload *LibpodBuildImageOKBody
}

func (o *LibpodBuildImageOK) Error() string {
	return fmt.Sprintf("[POST /libpod/build][%d] libpodBuildImageOK  %+v", 200, o.Payload)
}
func (o *LibpodBuildImageOK) GetPayload() *LibpodBuildImageOKBody {
	return o.Payload
}

func (o *LibpodBuildImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodBuildImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodBuildImageBadRequest creates a LibpodBuildImageBadRequest with default headers values
func NewLibpodBuildImageBadRequest() *LibpodBuildImageBadRequest {
	return &LibpodBuildImageBadRequest{}
}

/* LibpodBuildImageBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type LibpodBuildImageBadRequest struct {
	Payload *LibpodBuildImageBadRequestBody
}

func (o *LibpodBuildImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/build][%d] libpodBuildImageBadRequest  %+v", 400, o.Payload)
}
func (o *LibpodBuildImageBadRequest) GetPayload() *LibpodBuildImageBadRequestBody {
	return o.Payload
}

func (o *LibpodBuildImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodBuildImageBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodBuildImageInternalServerError creates a LibpodBuildImageInternalServerError with default headers values
func NewLibpodBuildImageInternalServerError() *LibpodBuildImageInternalServerError {
	return &LibpodBuildImageInternalServerError{}
}

/* LibpodBuildImageInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodBuildImageInternalServerError struct {
	Payload *LibpodBuildImageInternalServerErrorBody
}

func (o *LibpodBuildImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/build][%d] libpodBuildImageInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodBuildImageInternalServerError) GetPayload() *LibpodBuildImageInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodBuildImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodBuildImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodBuildImageBadRequestBody libpod build image bad request body
swagger:model LibpodBuildImageBadRequestBody
*/
type LibpodBuildImageBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod build image bad request body
func (o *LibpodBuildImageBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod build image bad request body based on context it is used
func (o *LibpodBuildImageBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodBuildImageBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodBuildImageBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LibpodBuildImageBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodBuildImageInternalServerErrorBody libpod build image internal server error body
swagger:model LibpodBuildImageInternalServerErrorBody
*/
type LibpodBuildImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod build image internal server error body
func (o *LibpodBuildImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod build image internal server error body based on context it is used
func (o *LibpodBuildImageInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodBuildImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodBuildImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodBuildImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodBuildImageOKBody libpod build image o k body
swagger:model LibpodBuildImageOKBody
*/
type LibpodBuildImageOKBody struct {

	// output from build process
	// Example: (build details...)\nSuccessfully built 8ba084515c724cbf90d447a63600c0a6\n
	// Required: true
	Stream *string `json:"stream"`
}

// Validate validates this libpod build image o k body
func (o *LibpodBuildImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStream(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodBuildImageOKBody) validateStream(formats strfmt.Registry) error {

	if err := validate.Required("libpodBuildImageOK"+"."+"stream", "body", o.Stream); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this libpod build image o k body based on context it is used
func (o *LibpodBuildImageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodBuildImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodBuildImageOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodBuildImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
