// Code generated by go-swagger; DO NOT EDIT.

package system

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
	"github.com/go-openapi/swag"
)

// NewLibpodGetEventsParams creates a new LibpodGetEventsParams object
// with the default values initialized.
func NewLibpodGetEventsParams() *LibpodGetEventsParams {
	var (
		streamDefault = bool(true)
	)
	return &LibpodGetEventsParams{
		Stream: &streamDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodGetEventsParamsWithTimeout creates a new LibpodGetEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodGetEventsParamsWithTimeout(timeout time.Duration) *LibpodGetEventsParams {
	var (
		streamDefault = bool(true)
	)
	return &LibpodGetEventsParams{
		Stream: &streamDefault,

		timeout: timeout,
	}
}

// NewLibpodGetEventsParamsWithContext creates a new LibpodGetEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodGetEventsParamsWithContext(ctx context.Context) *LibpodGetEventsParams {
	var (
		streamDefault = bool(true)
	)
	return &LibpodGetEventsParams{
		Stream: &streamDefault,

		Context: ctx,
	}
}

// NewLibpodGetEventsParamsWithHTTPClient creates a new LibpodGetEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodGetEventsParamsWithHTTPClient(client *http.Client) *LibpodGetEventsParams {
	var (
		streamDefault = bool(true)
	)
	return &LibpodGetEventsParams{
		Stream:     &streamDefault,
		HTTPClient: client,
	}
}

/*LibpodGetEventsParams contains all the parameters to send to the API endpoint
for the libpod get events operation typically these are written to a http.Request
*/
type LibpodGetEventsParams struct {

	/*Filters
	  JSON encoded map[string][]string of constraints

	*/
	Filters *string
	/*Since
	  start streaming events from this time

	*/
	Since *string
	/*Stream
	  when false, do not follow events

	*/
	Stream *bool
	/*Until
	  stop streaming events later than this

	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod get events params
func (o *LibpodGetEventsParams) WithTimeout(timeout time.Duration) *LibpodGetEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod get events params
func (o *LibpodGetEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod get events params
func (o *LibpodGetEventsParams) WithContext(ctx context.Context) *LibpodGetEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod get events params
func (o *LibpodGetEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod get events params
func (o *LibpodGetEventsParams) WithHTTPClient(client *http.Client) *LibpodGetEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod get events params
func (o *LibpodGetEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the libpod get events params
func (o *LibpodGetEventsParams) WithFilters(filters *string) *LibpodGetEventsParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the libpod get events params
func (o *LibpodGetEventsParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithSince adds the since to the libpod get events params
func (o *LibpodGetEventsParams) WithSince(since *string) *LibpodGetEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the libpod get events params
func (o *LibpodGetEventsParams) SetSince(since *string) {
	o.Since = since
}

// WithStream adds the stream to the libpod get events params
func (o *LibpodGetEventsParams) WithStream(stream *bool) *LibpodGetEventsParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the libpod get events params
func (o *LibpodGetEventsParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WithUntil adds the until to the libpod get events params
func (o *LibpodGetEventsParams) WithUntil(until *string) *LibpodGetEventsParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the libpod get events params
func (o *LibpodGetEventsParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodGetEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.Stream != nil {

		// query param stream
		var qrStream bool
		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {
			if err := r.SetQueryParam("stream", qStream); err != nil {
				return err
			}
		}

	}

	if o.Until != nil {

		// query param until
		var qrUntil string
		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {
			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}