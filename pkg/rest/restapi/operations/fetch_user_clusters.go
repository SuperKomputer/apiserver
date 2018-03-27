// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FetchUserClustersHandlerFunc turns a function with the right signature into a fetch user clusters handler
type FetchUserClustersHandlerFunc func(FetchUserClustersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchUserClustersHandlerFunc) Handle(params FetchUserClustersParams) middleware.Responder {
	return fn(params)
}

// FetchUserClustersHandler interface for that can handle valid fetch user clusters params
type FetchUserClustersHandler interface {
	Handle(FetchUserClustersParams) middleware.Responder
}

// NewFetchUserClusters creates a new http.Handler for the fetch user clusters operation
func NewFetchUserClusters(ctx *middleware.Context, handler FetchUserClustersHandler) *FetchUserClusters {
	return &FetchUserClusters{Context: ctx, Handler: handler}
}

/*FetchUserClusters swagger:route GET /users/{username}/clusters fetchUserClusters

fetch the list of using clusters of user with given username

*/
type FetchUserClusters struct {
	Context *middleware.Context
	Handler FetchUserClustersHandler
}

func (o *FetchUserClusters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFetchUserClustersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
