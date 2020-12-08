// Code generated by go-swagger; DO NOT EDIT.

package rbds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbdsHandlerFunc turns a function with the right signature into a list rbds handler
type ListRbdsHandlerFunc func(ListRbdsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbdsHandlerFunc) Handle(params ListRbdsParams) middleware.Responder {
	return fn(params)
}

// ListRbdsHandler interface for that can handle valid list rbds params
type ListRbdsHandler interface {
	Handle(ListRbdsParams) middleware.Responder
}

// NewListRbds creates a new http.Handler for the list rbds operation
func NewListRbds(ctx *middleware.Context, handler ListRbdsHandler) *ListRbds {
	return &ListRbds{Context: ctx, Handler: handler}
}

/*ListRbds swagger:route GET /rbd rbds listRbds

ListRbds list rbds API

*/
type ListRbds struct {
	Context *middleware.Context
	Handler ListRbdsHandler
}

func (o *ListRbds) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbdsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
