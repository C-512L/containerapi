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

// LibpodPruneContainersReader is a Reader for the LibpodPruneContainers structure.
type LibpodPruneContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodPruneContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodPruneContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodPruneContainersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodPruneContainersOK creates a LibpodPruneContainersOK with default headers values
func NewLibpodPruneContainersOK() *LibpodPruneContainersOK {
	return &LibpodPruneContainersOK{}
}

/*LibpodPruneContainersOK handles this case with default header values.

Prune containers
*/
type LibpodPruneContainersOK struct {
	Payload []*LibpodPruneContainersOKBodyItems0
}

func (o *LibpodPruneContainersOK) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/prune][%d] libpodPruneContainersOK  %+v", 200, o.Payload)
}

func (o *LibpodPruneContainersOK) GetPayload() []*LibpodPruneContainersOKBodyItems0 {
	return o.Payload
}

func (o *LibpodPruneContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodPruneContainersInternalServerError creates a LibpodPruneContainersInternalServerError with default headers values
func NewLibpodPruneContainersInternalServerError() *LibpodPruneContainersInternalServerError {
	return &LibpodPruneContainersInternalServerError{}
}

/*LibpodPruneContainersInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodPruneContainersInternalServerError struct {
	Payload *LibpodPruneContainersInternalServerErrorBody
}

func (o *LibpodPruneContainersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/prune][%d] libpodPruneContainersInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodPruneContainersInternalServerError) GetPayload() *LibpodPruneContainersInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodPruneContainersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodPruneContainersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodPruneContainersInternalServerErrorBody libpod prune containers internal server error body
swagger:model LibpodPruneContainersInternalServerErrorBody
*/
type LibpodPruneContainersInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod prune containers internal server error body
func (o *LibpodPruneContainersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPruneContainersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPruneContainersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodPruneContainersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodPruneContainersOKBodyItems0 libpod prune containers o k body items0
swagger:model LibpodPruneContainersOKBodyItems0
*/
type LibpodPruneContainersOKBodyItems0 struct {

	// prune error
	PruneError string `json:"error,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// space reclaimed
	SpaceReclaimed int64 `json:"space,omitempty"`
}

// Validate validates this libpod prune containers o k body items0
func (o *LibpodPruneContainersOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPruneContainersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPruneContainersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res LibpodPruneContainersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}