// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetMountOverlayURL generates an URL for the get mount overlay operation
type GetMountOverlayURL struct {
	Lower int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMountOverlayURL) WithBasePath(bp string) *GetMountOverlayURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMountOverlayURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMountOverlayURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/mount/overlay/{lower}"

	lower := swag.FormatInt64(o.Lower)
	if lower != "" {
		_path = strings.Replace(_path, "{lower}", lower, -1)
	} else {
		return nil, errors.New("lower is required on GetMountOverlayURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetMountOverlayURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMountOverlayURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMountOverlayURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMountOverlayURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMountOverlayURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetMountOverlayURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
