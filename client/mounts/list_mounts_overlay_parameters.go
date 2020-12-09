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

// NewListMountsOverlayParams creates a new ListMountsOverlayParams object
// with the default values initialized.
func NewListMountsOverlayParams() *ListMountsOverlayParams {

	return &ListMountsOverlayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMountsOverlayParamsWithTimeout creates a new ListMountsOverlayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMountsOverlayParamsWithTimeout(timeout time.Duration) *ListMountsOverlayParams {

	return &ListMountsOverlayParams{

		timeout: timeout,
	}
}

// NewListMountsOverlayParamsWithContext creates a new ListMountsOverlayParams object
// with the default values initialized, and the ability to set a context for a request
func NewListMountsOverlayParamsWithContext(ctx context.Context) *ListMountsOverlayParams {

	return &ListMountsOverlayParams{

		Context: ctx,
	}
}

// NewListMountsOverlayParamsWithHTTPClient creates a new ListMountsOverlayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMountsOverlayParamsWithHTTPClient(client *http.Client) *ListMountsOverlayParams {

	return &ListMountsOverlayParams{
		HTTPClient: client,
	}
}

/*ListMountsOverlayParams contains all the parameters to send to the API endpoint
for the list mounts overlay operation typically these are written to a http.Request
*/
type ListMountsOverlayParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list mounts overlay params
func (o *ListMountsOverlayParams) WithTimeout(timeout time.Duration) *ListMountsOverlayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mounts overlay params
func (o *ListMountsOverlayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mounts overlay params
func (o *ListMountsOverlayParams) WithContext(ctx context.Context) *ListMountsOverlayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mounts overlay params
func (o *ListMountsOverlayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mounts overlay params
func (o *ListMountsOverlayParams) WithHTTPClient(client *http.Client) *ListMountsOverlayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mounts overlay params
func (o *ListMountsOverlayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMountsOverlayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}