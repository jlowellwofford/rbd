// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateContainerHandlerFunc turns a function with the right signature into a create container handler
type CreateContainerHandlerFunc func(CreateContainerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateContainerHandlerFunc) Handle(params CreateContainerParams) middleware.Responder {
	return fn(params)
}

// CreateContainerHandler interface for that can handle valid create container params
type CreateContainerHandler interface {
	Handle(CreateContainerParams) middleware.Responder
}

// NewCreateContainer creates a new http.Handler for the create container operation
func NewCreateContainer(ctx *middleware.Context, handler CreateContainerHandler) *CreateContainer {
	return &CreateContainer{Context: ctx, Handler: handler}
}

/*CreateContainer swagger:route POST /container containers createContainer

Create a container

*/
type CreateContainer struct {
	Context *middleware.Context
	Handler CreateContainerHandler
}

func (o *CreateContainer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateContainerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}