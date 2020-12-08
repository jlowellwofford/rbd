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
)

// NewListMountsRbdParams creates a new ListMountsRbdParams object
// with the default values initialized.
func NewListMountsRbdParams() *ListMountsRbdParams {

	return &ListMountsRbdParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMountsRbdParamsWithTimeout creates a new ListMountsRbdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMountsRbdParamsWithTimeout(timeout time.Duration) *ListMountsRbdParams {

	return &ListMountsRbdParams{

		timeout: timeout,
	}
}

// NewListMountsRbdParamsWithContext creates a new ListMountsRbdParams object
// with the default values initialized, and the ability to set a context for a request
func NewListMountsRbdParamsWithContext(ctx context.Context) *ListMountsRbdParams {

	return &ListMountsRbdParams{

		Context: ctx,
	}
}

// NewListMountsRbdParamsWithHTTPClient creates a new ListMountsRbdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMountsRbdParamsWithHTTPClient(client *http.Client) *ListMountsRbdParams {

	return &ListMountsRbdParams{
		HTTPClient: client,
	}
}

/*ListMountsRbdParams contains all the parameters to send to the API endpoint
for the list mounts rbd operation typically these are written to a http.Request
*/
type ListMountsRbdParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list mounts rbd params
func (o *ListMountsRbdParams) WithTimeout(timeout time.Duration) *ListMountsRbdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mounts rbd params
func (o *ListMountsRbdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mounts rbd params
func (o *ListMountsRbdParams) WithContext(ctx context.Context) *ListMountsRbdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mounts rbd params
func (o *ListMountsRbdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mounts rbd params
func (o *ListMountsRbdParams) WithHTTPClient(client *http.Client) *ListMountsRbdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mounts rbd params
func (o *ListMountsRbdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMountsRbdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
