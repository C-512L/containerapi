// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewLibpodMountContainerParams creates a new LibpodMountContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodMountContainerParams() *LibpodMountContainerParams {
	return &LibpodMountContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodMountContainerParamsWithTimeout creates a new LibpodMountContainerParams object
// with the ability to set a timeout on a request.
func NewLibpodMountContainerParamsWithTimeout(timeout time.Duration) *LibpodMountContainerParams {
	return &LibpodMountContainerParams{
		timeout: timeout,
	}
}

// NewLibpodMountContainerParamsWithContext creates a new LibpodMountContainerParams object
// with the ability to set a context for a request.
func NewLibpodMountContainerParamsWithContext(ctx context.Context) *LibpodMountContainerParams {
	return &LibpodMountContainerParams{
		Context: ctx,
	}
}

// NewLibpodMountContainerParamsWithHTTPClient creates a new LibpodMountContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodMountContainerParamsWithHTTPClient(client *http.Client) *LibpodMountContainerParams {
	return &LibpodMountContainerParams{
		HTTPClient: client,
	}
}

/* LibpodMountContainerParams contains all the parameters to send to the API endpoint
   for the libpod mount container operation.

   Typically these are written to a http.Request.
*/
type LibpodMountContainerParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod mount container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodMountContainerParams) WithDefaults() *LibpodMountContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod mount container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodMountContainerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the libpod mount container params
func (o *LibpodMountContainerParams) WithTimeout(timeout time.Duration) *LibpodMountContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod mount container params
func (o *LibpodMountContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod mount container params
func (o *LibpodMountContainerParams) WithContext(ctx context.Context) *LibpodMountContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod mount container params
func (o *LibpodMountContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod mount container params
func (o *LibpodMountContainerParams) WithHTTPClient(client *http.Client) *LibpodMountContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod mount container params
func (o *LibpodMountContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod mount container params
func (o *LibpodMountContainerParams) WithName(name string) *LibpodMountContainerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod mount container params
func (o *LibpodMountContainerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodMountContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
