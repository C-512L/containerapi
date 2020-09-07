// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodGetEventsReader is a Reader for the LibpodGetEvents structure.
type LibpodGetEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodGetEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodGetEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodGetEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodGetEventsOK creates a LibpodGetEventsOK with default headers values
func NewLibpodGetEventsOK() *LibpodGetEventsOK {
	return &LibpodGetEventsOK{}
}

/*LibpodGetEventsOK handles this case with default header values.

returns a string of json data describing an event
*/
type LibpodGetEventsOK struct {
}

func (o *LibpodGetEventsOK) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] libpodGetEventsOK ", 200)
}

func (o *LibpodGetEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodGetEventsInternalServerError creates a LibpodGetEventsInternalServerError with default headers values
func NewLibpodGetEventsInternalServerError() *LibpodGetEventsInternalServerError {
	return &LibpodGetEventsInternalServerError{}
}

/*LibpodGetEventsInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodGetEventsInternalServerError struct {
	Payload *LibpodGetEventsInternalServerErrorBody
}

func (o *LibpodGetEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] libpodGetEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodGetEventsInternalServerError) GetPayload() *LibpodGetEventsInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodGetEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodGetEventsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodGetEventsInternalServerErrorBody libpod get events internal server error body
swagger:model LibpodGetEventsInternalServerErrorBody
*/
type LibpodGetEventsInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod get events internal server error body
func (o *LibpodGetEventsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetEventsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetEventsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodGetEventsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
