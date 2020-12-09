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

// NewCreateContainerParams creates a new CreateContainerParams object
// with the default values initialized.
func NewCreateContainerParams() *CreateContainerParams {

	return &CreateContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateContainerParamsWithTimeout creates a new CreateContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateContainerParamsWithTimeout(timeout time.Duration) *CreateContainerParams {

	return &CreateContainerParams{

		timeout: timeout,
	}
}

// NewCreateContainerParamsWithContext creates a new CreateContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateContainerParamsWithContext(ctx context.Context) *CreateContainerParams {

	return &CreateContainerParams{

		Context: ctx,
	}
}

// NewCreateContainerParamsWithHTTPClient creates a new CreateContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateContainerParamsWithHTTPClient(client *http.Client) *CreateContainerParams {

	return &CreateContainerParams{
		HTTPClient: client,
	}
}

/*CreateContainerParams contains all the parameters to send to the API endpoint
for the create container operation typically these are written to a http.Request
*/
type CreateContainerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create container params
func (o *CreateContainerParams) WithTimeout(timeout time.Duration) *CreateContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create container params
func (o *CreateContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create container params
func (o *CreateContainerParams) WithContext(ctx context.Context) *CreateContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create container params
func (o *CreateContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create container params
func (o *CreateContainerParams) WithHTTPClient(client *http.Client) *CreateContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create container params
func (o *CreateContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}