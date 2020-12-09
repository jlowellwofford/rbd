// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetMountOverlay(params *GetMountOverlayParams) (*GetMountOverlayOK, error)

	GetMountRbd(params *GetMountRbdParams) (*GetMountRbdOK, error)

	ListMountsOverlay(params *ListMountsOverlayParams) (*ListMountsOverlayOK, error)

	ListMountsRbd(params *ListMountsRbdParams) (*ListMountsRbdOK, error)

	MountOverlay(params *MountOverlayParams) (*MountOverlayCreated, error)

	MountRbd(params *MountRbdParams) (*MountRbdCreated, error)

	UnmountOverlay(params *UnmountOverlayParams) (*UnmountOverlayNoContent, error)

	UnmountRbd(params *UnmountRbdParams) (*UnmountRbdNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetMountOverlay get mount overlay API
*/
func (a *Client) GetMountOverlay(params *GetMountOverlayParams) (*GetMountOverlayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMountOverlayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_mount_overlay",
		Method:             "GET",
		PathPattern:        "/mount/overlay/{id}",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMountOverlayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMountOverlayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMountOverlayDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetMountRbd get mount rbd API
*/
func (a *Client) GetMountRbd(params *GetMountRbdParams) (*GetMountRbdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMountRbdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_mount_rbd",
		Method:             "GET",
		PathPattern:        "/mount/rbd/{id}",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMountRbdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMountRbdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMountRbdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListMountsOverlay list mounts overlay API
*/
func (a *Client) ListMountsOverlay(params *ListMountsOverlayParams) (*ListMountsOverlayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMountsOverlayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_mounts_overlay",
		Method:             "GET",
		PathPattern:        "/mount/overlay",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMountsOverlayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMountsOverlayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListMountsOverlayDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListMountsRbd list mounts rbd API
*/
func (a *Client) ListMountsRbd(params *ListMountsRbdParams) (*ListMountsRbdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMountsRbdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_mounts_rbd",
		Method:             "GET",
		PathPattern:        "/mount/rbd",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMountsRbdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMountsRbdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListMountsRbdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MountOverlay mount overlay API
*/
func (a *Client) MountOverlay(params *MountOverlayParams) (*MountOverlayCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMountOverlayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mount_overlay",
		Method:             "POST",
		PathPattern:        "/mount/overlay",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MountOverlayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MountOverlayCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MountOverlayDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MountRbd mount rbd API
*/
func (a *Client) MountRbd(params *MountRbdParams) (*MountRbdCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMountRbdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mount_rbd",
		Method:             "POST",
		PathPattern:        "/mount/rbd",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MountRbdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MountRbdCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MountRbdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UnmountOverlay unmount overlay API
*/
func (a *Client) UnmountOverlay(params *UnmountOverlayParams) (*UnmountOverlayNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnmountOverlayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unmount_overlay",
		Method:             "DELETE",
		PathPattern:        "/mount/overlay/{id}",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnmountOverlayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnmountOverlayNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnmountOverlayDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UnmountRbd unmount rbd API
*/
func (a *Client) UnmountRbd(params *UnmountRbdParams) (*UnmountRbdNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnmountRbdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unmount_rbd",
		Method:             "DELETE",
		PathPattern:        "/mount/rbd/{id}",
		ProducesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		ConsumesMediaTypes: []string{"application/github.com.bensallen.rbd.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnmountRbdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnmountRbdNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnmountRbdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
