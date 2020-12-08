// Code generated by go-swagger; DO NOT EDIT.

package rbds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UnmapRbdHandlerFunc turns a function with the right signature into a unmap rbd handler
type UnmapRbdHandlerFunc func(UnmapRbdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UnmapRbdHandlerFunc) Handle(params UnmapRbdParams) middleware.Responder {
	return fn(params)
}

// UnmapRbdHandler interface for that can handle valid unmap rbd params
type UnmapRbdHandler interface {
	Handle(UnmapRbdParams) middleware.Responder
}

// NewUnmapRbd creates a new http.Handler for the unmap rbd operation
func NewUnmapRbd(ctx *middleware.Context, handler UnmapRbdHandler) *UnmapRbd {
	return &UnmapRbd{Context: ctx, Handler: handler}
}

/*UnmapRbd swagger:route DELETE /rbd/{id} rbds unmapRbd

UnmapRbd unmap rbd API

*/
type UnmapRbd struct {
	Context *middleware.Context
	Handler UnmapRbdHandler
}

func (o *UnmapRbd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUnmapRbdParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
