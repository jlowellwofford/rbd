// Code generated by go-swagger; DO NOT EDIT.

package mounts

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

// NewGetMountRbdParams creates a new GetMountRbdParams object
// with the default values initialized.
func NewGetMountRbdParams() *GetMountRbdParams {
	var ()
	return &GetMountRbdParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMountRbdParamsWithTimeout creates a new GetMountRbdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMountRbdParamsWithTimeout(timeout time.Duration) *GetMountRbdParams {
	var ()
	return &GetMountRbdParams{

		timeout: timeout,
	}
}

// NewGetMountRbdParamsWithContext creates a new GetMountRbdParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMountRbdParamsWithContext(ctx context.Context) *GetMountRbdParams {
	var ()
	return &GetMountRbdParams{

		Context: ctx,
	}
}

// NewGetMountRbdParamsWithHTTPClient creates a new GetMountRbdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMountRbdParamsWithHTTPClient(client *http.Client) *GetMountRbdParams {
	var ()
	return &GetMountRbdParams{
		HTTPClient: client,
	}
}

/*GetMountRbdParams contains all the parameters to send to the API endpoint
for the get mount rbd operation typically these are written to a http.Request
*/
type GetMountRbdParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mount rbd params
func (o *GetMountRbdParams) WithTimeout(timeout time.Duration) *GetMountRbdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mount rbd params
func (o *GetMountRbdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mount rbd params
func (o *GetMountRbdParams) WithContext(ctx context.Context) *GetMountRbdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mount rbd params
func (o *GetMountRbdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mount rbd params
func (o *GetMountRbdParams) WithHTTPClient(client *http.Client) *GetMountRbdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mount rbd params
func (o *GetMountRbdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get mount rbd params
func (o *GetMountRbdParams) WithID(id int64) *GetMountRbdParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get mount rbd params
func (o *GetMountRbdParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMountRbdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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