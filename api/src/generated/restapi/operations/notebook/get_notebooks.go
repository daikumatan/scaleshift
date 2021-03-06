// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNotebooksHandlerFunc turns a function with the right signature into a get notebooks handler
type GetNotebooksHandlerFunc func(GetNotebooksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNotebooksHandlerFunc) Handle(params GetNotebooksParams) middleware.Responder {
	return fn(params)
}

// GetNotebooksHandler interface for that can handle valid get notebooks params
type GetNotebooksHandler interface {
	Handle(GetNotebooksParams) middleware.Responder
}

// NewGetNotebooks creates a new http.Handler for the get notebooks operation
func NewGetNotebooks(ctx *middleware.Context, handler GetNotebooksHandler) *GetNotebooks {
	return &GetNotebooks{Context: ctx, Handler: handler}
}

/*GetNotebooks swagger:route GET /notebooks notebook getNotebooks

returns Jupyter notebook information


*/
type GetNotebooks struct {
	Context *middleware.Context
	Handler GetNotebooksHandler
}

func (o *GetNotebooks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNotebooksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
