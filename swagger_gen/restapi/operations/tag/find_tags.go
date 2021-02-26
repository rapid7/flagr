// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindTagsHandlerFunc turns a function with the right signature into a find tags handler
type FindTagsHandlerFunc func(FindTagsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindTagsHandlerFunc) Handle(params FindTagsParams) middleware.Responder {
	return fn(params)
}

// FindTagsHandler interface for that can handle valid find tags params
type FindTagsHandler interface {
	Handle(FindTagsParams) middleware.Responder
}

// NewFindTags creates a new http.Handler for the find tags operation
func NewFindTags(ctx *middleware.Context, handler FindTagsHandler) *FindTags {
	return &FindTags{Context: ctx, Handler: handler}
}

/* FindTags swagger:route GET /flags/{flagID}/tags tag findTags

FindTags find tags API

*/
type FindTags struct {
	Context *middleware.Context
	Handler FindTagsHandler
}

func (o *FindTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindTagsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}