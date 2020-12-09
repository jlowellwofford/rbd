// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUnmountOverlayParams creates a new UnmountOverlayParams object
// no default values defined in spec.
func NewUnmountOverlayParams() UnmountOverlayParams {

	return UnmountOverlayParams{}
}

// UnmountOverlayParams contains all the bound params for the unmount overlay operation
// typically these are obtained from a http.Request
//
// swagger:parameters unmount_overlay
type UnmountOverlayParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Lower int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUnmountOverlayParams() beforehand.
func (o *UnmountOverlayParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rLower, rhkLower, _ := route.Params.GetOK("lower")
	if err := o.bindLower(rLower, rhkLower, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLower binds and validates parameter Lower from path.
func (o *UnmountOverlayParams) bindLower(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("lower", "path", "int64", raw)
	}
	o.Lower = value

	return nil
}
