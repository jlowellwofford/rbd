// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDeleteContainerParams creates a new DeleteContainerParams object
// with the default values initialized.
func NewDeleteContainerParams() *DeleteContainerParams {
	var ()
	return &DeleteContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContainerParamsWithTimeout creates a new DeleteContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteContainerParamsWithTimeout(timeout time.Duration) *DeleteContainerParams {
	var ()
	return &DeleteContainerParams{

		timeout: timeout,
	}
}

// NewDeleteContainerParamsWithContext creates a new DeleteContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteContainerParamsWithContext(ctx context.Context) *DeleteContainerParams {
	var ()
	return &DeleteContainerParams{

		Context: ctx,
	}
}

// NewDeleteContainerParamsWithHTTPClient creates a new DeleteContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteContainerParamsWithHTTPClient(client *http.Client) *DeleteContainerParams {
	var ()
	return &DeleteContainerParams{
		HTTPClient: client,
	}
}

/*DeleteContainerParams contains all the parameters to send to the API endpoint
for the delete container operation typically these are written to a http.Request
*/
type DeleteContainerParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete container params
func (o *DeleteContainerParams) WithTimeout(timeout time.Duration) *DeleteContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete container params
func (o *DeleteContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete container params
func (o *DeleteContainerParams) WithContext(ctx context.Context) *DeleteContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete container params
func (o *DeleteContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete container params
func (o *DeleteContainerParams) WithHTTPClient(client *http.Client) *DeleteContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete container params
func (o *DeleteContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete container params
func (o *DeleteContainerParams) WithID(id int64) *DeleteContainerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete container params
func (o *DeleteContainerParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
