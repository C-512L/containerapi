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

// LibpodStatsContainerReader is a Reader for the LibpodStatsContainer structure.
type LibpodStatsContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodStatsContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodStatsContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodStatsContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodStatsContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodStatsContainerOK creates a LibpodStatsContainerOK with default headers values
func NewLibpodStatsContainerOK() *LibpodStatsContainerOK {
	return &LibpodStatsContainerOK{}
}

/*LibpodStatsContainerOK handles this case with default header values.

no error
*/
type LibpodStatsContainerOK struct {
}

func (o *LibpodStatsContainerOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] libpodStatsContainerOK ", 200)
}

func (o *LibpodStatsContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodStatsContainerNotFound creates a LibpodStatsContainerNotFound with default headers values
func NewLibpodStatsContainerNotFound() *LibpodStatsContainerNotFound {
	return &LibpodStatsContainerNotFound{}
}

/*LibpodStatsContainerNotFound handles this case with default header values.

No such container
*/
type LibpodStatsContainerNotFound struct {
	Payload *LibpodStatsContainerNotFoundBody
}

func (o *LibpodStatsContainerNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] libpodStatsContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodStatsContainerNotFound) GetPayload() *LibpodStatsContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodStatsContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStatsContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodStatsContainerInternalServerError creates a LibpodStatsContainerInternalServerError with default headers values
func NewLibpodStatsContainerInternalServerError() *LibpodStatsContainerInternalServerError {
	return &LibpodStatsContainerInternalServerError{}
}

/*LibpodStatsContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodStatsContainerInternalServerError struct {
	Payload *LibpodStatsContainerInternalServerErrorBody
}

func (o *LibpodStatsContainerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] libpodStatsContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodStatsContainerInternalServerError) GetPayload() *LibpodStatsContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodStatsContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodStatsContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodStatsContainerInternalServerErrorBody libpod stats container internal server error body
swagger:model LibpodStatsContainerInternalServerErrorBody
*/
type LibpodStatsContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stats container internal server error body
func (o *LibpodStatsContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStatsContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStatsContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodStatsContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodStatsContainerNotFoundBody libpod stats container not found body
swagger:model LibpodStatsContainerNotFoundBody
*/
type LibpodStatsContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod stats container not found body
func (o *LibpodStatsContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodStatsContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodStatsContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodStatsContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
