// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateUserHandlerFunc turns a function with the right signature into a update user handler
type UpdateUserHandlerFunc func(UpdateUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserHandlerFunc) Handle(params UpdateUserParams) middleware.Responder {
	return fn(params)
}

// UpdateUserHandler interface for that can handle valid update user params
type UpdateUserHandler interface {
	Handle(UpdateUserParams) middleware.Responder
}

// NewUpdateUser creates a new http.Handler for the update user operation
func NewUpdateUser(ctx *middleware.Context, handler UpdateUserHandler) *UpdateUser {
	return &UpdateUser{Context: ctx, Handler: handler}
}

/*UpdateUser swagger:route PUT /users/{username} updateUser

updates a user

updates a user with the given update config

*/
type UpdateUser struct {
	Context *middleware.Context
	Handler UpdateUserHandler
}

func (o *UpdateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
