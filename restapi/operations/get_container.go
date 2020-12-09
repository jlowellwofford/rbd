// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetContainerHandlerFunc turns a function with the right signature into a get container handler
type GetContainerHandlerFunc func(GetContainerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContainerHandlerFunc) Handle(params GetContainerParams) middleware.Responder {
	return fn(params)
}

// GetContainerHandler interface for that can handle valid get container params
type GetContainerHandler interface {
	Handle(GetContainerParams) middleware.Responder
}

// NewGetContainer creates a new http.Handler for the get container operation
func NewGetContainer(ctx *middleware.Context, handler GetContainerHandler) *GetContainer {
	return &GetContainer{Context: ctx, Handler: handler}
}

/*GetContainer swagger:route GET /container/{id} getContainer

Get a container definition

*/
type GetContainer struct {
	Context *middleware.Context
	Handler GetContainerHandler
}

func (o *GetContainer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetContainerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
