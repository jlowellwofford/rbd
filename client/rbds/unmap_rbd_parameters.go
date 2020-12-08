// Code generated by go-swagger; DO NOT EDIT.

package rbds

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

// NewUnmapRbdParams creates a new UnmapRbdParams object
// with the default values initialized.
func NewUnmapRbdParams() *UnmapRbdParams {
	var ()
	return &UnmapRbdParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnmapRbdParamsWithTimeout creates a new UnmapRbdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnmapRbdParamsWithTimeout(timeout time.Duration) *UnmapRbdParams {
	var ()
	return &UnmapRbdParams{

		timeout: timeout,
	}
}

// NewUnmapRbdParamsWithContext creates a new UnmapRbdParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnmapRbdParamsWithContext(ctx context.Context) *UnmapRbdParams {
	var ()
	return &UnmapRbdParams{

		Context: ctx,
	}
}

// NewUnmapRbdParamsWithHTTPClient creates a new UnmapRbdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnmapRbdParamsWithHTTPClient(client *http.Client) *UnmapRbdParams {
	var ()
	return &UnmapRbdParams{
		HTTPClient: client,
	}
}

/*UnmapRbdParams contains all the parameters to send to the API endpoint
for the unmap rbd operation typically these are written to a http.Request
*/
type UnmapRbdParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unmap rbd params
func (o *UnmapRbdParams) WithTimeout(timeout time.Duration) *UnmapRbdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unmap rbd params
func (o *UnmapRbdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unmap rbd params
func (o *UnmapRbdParams) WithContext(ctx context.Context) *UnmapRbdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unmap rbd params
func (o *UnmapRbdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unmap rbd params
func (o *UnmapRbdParams) WithHTTPClient(client *http.Client) *UnmapRbdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unmap rbd params
func (o *UnmapRbdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unmap rbd params
func (o *UnmapRbdParams) WithID(id int64) *UnmapRbdParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unmap rbd params
func (o *UnmapRbdParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnmapRbdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
