// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2CancelInstallationHandlerFunc turns a function with the right signature into a v2 cancel installation handler
type V2CancelInstallationHandlerFunc func(V2CancelInstallationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2CancelInstallationHandlerFunc) Handle(params V2CancelInstallationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2CancelInstallationHandler interface for that can handle valid v2 cancel installation params
type V2CancelInstallationHandler interface {
	Handle(V2CancelInstallationParams, interface{}) middleware.Responder
}

// NewV2CancelInstallation creates a new http.Handler for the v2 cancel installation operation
func NewV2CancelInstallation(ctx *middleware.Context, handler V2CancelInstallationHandler) *V2CancelInstallation {
	return &V2CancelInstallation{Context: ctx, Handler: handler}
}

/*V2CancelInstallation swagger:route POST /v2/clusters/{cluster_id}/actions/cancel installer v2CancelInstallation

Cancels an ongoing installation.

*/
type V2CancelInstallation struct {
	Context *middleware.Context
	Handler V2CancelInstallationHandler
}

func (o *V2CancelInstallation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV2CancelInstallationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
