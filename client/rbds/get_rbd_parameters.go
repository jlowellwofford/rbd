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

// NewGetRbdParams creates a new GetRbdParams object
// with the default values initialized.
func NewGetRbdParams() *GetRbdParams {
	var ()
	return &GetRbdParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRbdParamsWithTimeout creates a new GetRbdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRbdParamsWithTimeout(timeout time.Duration) *GetRbdParams {
	var ()
	return &GetRbdParams{

		timeout: timeout,
	}
}

// NewGetRbdParamsWithContext creates a new GetRbdParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRbdParamsWithContext(ctx context.Context) *GetRbdParams {
	var ()
	return &GetRbdParams{

		Context: ctx,
	}
}

// NewGetRbdParamsWithHTTPClient creates a new GetRbdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRbdParamsWithHTTPClient(client *http.Client) *GetRbdParams {
	var ()
	return &GetRbdParams{
		HTTPClient: client,
	}
}

/*GetRbdParams contains all the parameters to send to the API endpoint
for the get rbd operation typically these are written to a http.Request
*/
type GetRbdParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rbd params
func (o *GetRbdParams) WithTimeout(timeout time.Duration) *GetRbdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rbd params
func (o *GetRbdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rbd params
func (o *GetRbdParams) WithContext(ctx context.Context) *GetRbdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rbd params
func (o *GetRbdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rbd params
func (o *GetRbdParams) WithHTTPClient(client *http.Client) *GetRbdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rbd params
func (o *GetRbdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get rbd params
func (o *GetRbdParams) WithID(id int64) *GetRbdParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get rbd params
func (o *GetRbdParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRbdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
