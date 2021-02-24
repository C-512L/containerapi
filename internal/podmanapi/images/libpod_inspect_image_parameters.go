// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewLibpodInspectImageParams creates a new LibpodInspectImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodInspectImageParams() *LibpodInspectImageParams {
	return &LibpodInspectImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodInspectImageParamsWithTimeout creates a new LibpodInspectImageParams object
// with the ability to set a timeout on a request.
func NewLibpodInspectImageParamsWithTimeout(timeout time.Duration) *LibpodInspectImageParams {
	return &LibpodInspectImageParams{
		timeout: timeout,
	}
}

// NewLibpodInspectImageParamsWithContext creates a new LibpodInspectImageParams object
// with the ability to set a context for a request.
func NewLibpodInspectImageParamsWithContext(ctx context.Context) *LibpodInspectImageParams {
	return &LibpodInspectImageParams{
		Context: ctx,
	}
}

// NewLibpodInspectImageParamsWithHTTPClient creates a new LibpodInspectImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodInspectImageParamsWithHTTPClient(client *http.Client) *LibpodInspectImageParams {
	return &LibpodInspectImageParams{
		HTTPClient: client,
	}
}

/* LibpodInspectImageParams contains all the parameters to send to the API endpoint
   for the libpod inspect image operation.

   Typically these are written to a http.Request.
*/
type LibpodInspectImageParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod inspect image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodInspectImageParams) WithDefaults() *LibpodInspectImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod inspect image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodInspectImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the libpod inspect image params
func (o *LibpodInspectImageParams) WithTimeout(timeout time.Duration) *LibpodInspectImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod inspect image params
func (o *LibpodInspectImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod inspect image params
func (o *LibpodInspectImageParams) WithContext(ctx context.Context) *LibpodInspectImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod inspect image params
func (o *LibpodInspectImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod inspect image params
func (o *LibpodInspectImageParams) WithHTTPClient(client *http.Client) *LibpodInspectImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod inspect image params
func (o *LibpodInspectImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod inspect image params
func (o *LibpodInspectImageParams) WithName(name string) *LibpodInspectImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod inspect image params
func (o *LibpodInspectImageParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodInspectImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name:.*
	if err := r.SetPathParam("name:.*", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
