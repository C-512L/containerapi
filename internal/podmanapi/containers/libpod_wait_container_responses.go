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

// LibpodWaitContainerReader is a Reader for the LibpodWaitContainer structure.
type LibpodWaitContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodWaitContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodWaitContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodWaitContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodWaitContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodWaitContainerOK creates a LibpodWaitContainerOK with default headers values
func NewLibpodWaitContainerOK() *LibpodWaitContainerOK {
	return &LibpodWaitContainerOK{}
}

/*LibpodWaitContainerOK handles this case with default header values.

exit code
*/
type LibpodWaitContainerOK struct {
	Payload string
}

func (o *LibpodWaitContainerOK) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] libpodWaitContainerOK  %+v", 200, o.Payload)
}

func (o *LibpodWaitContainerOK) GetPayload() string {
	return o.Payload
}

func (o *LibpodWaitContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodWaitContainerNotFound creates a LibpodWaitContainerNotFound with default headers values
func NewLibpodWaitContainerNotFound() *LibpodWaitContainerNotFound {
	return &LibpodWaitContainerNotFound{}
}

/*LibpodWaitContainerNotFound handles this case with default header values.

No such container
*/
type LibpodWaitContainerNotFound struct {
	Payload *LibpodWaitContainerNotFoundBody
}

func (o *LibpodWaitContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] libpodWaitContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodWaitContainerNotFound) GetPayload() *LibpodWaitContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodWaitContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodWaitContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodWaitContainerInternalServerError creates a LibpodWaitContainerInternalServerError with default headers values
func NewLibpodWaitContainerInternalServerError() *LibpodWaitContainerInternalServerError {
	return &LibpodWaitContainerInternalServerError{}
}

/*LibpodWaitContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodWaitContainerInternalServerError struct {
	Payload *LibpodWaitContainerInternalServerErrorBody
}

func (o *LibpodWaitContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/wait][%d] libpodWaitContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodWaitContainerInternalServerError) GetPayload() *LibpodWaitContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodWaitContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodWaitContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodWaitContainerInternalServerErrorBody libpod wait container internal server error body
swagger:model LibpodWaitContainerInternalServerErrorBody
*/
type LibpodWaitContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod wait container internal server error body
func (o *LibpodWaitContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodWaitContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodWaitContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodWaitContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodWaitContainerNotFoundBody libpod wait container not found body
swagger:model LibpodWaitContainerNotFoundBody
*/
type LibpodWaitContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod wait container not found body
func (o *LibpodWaitContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodWaitContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodWaitContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodWaitContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}