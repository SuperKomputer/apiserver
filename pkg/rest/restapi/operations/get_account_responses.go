// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/fanzhangio/superkomputer/pkg/rest/models"
)

// GetAccountOKCode is the HTTP code returned for type GetAccountOK
const GetAccountOKCode int = 200

/*GetAccountOK 200 response with account information

swagger:response getAccountOK
*/
type GetAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewGetAccountOK creates GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {

	return &GetAccountOK{}
}

// WithPayload adds the payload to the get account o k response
func (o *GetAccountOK) WithPayload(payload *models.Account) *GetAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account o k response
func (o *GetAccountOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountNotFoundCode is the HTTP code returned for type GetAccountNotFound
const GetAccountNotFoundCode int = 404

/*GetAccountNotFound The account was not found

swagger:response getAccountNotFound
*/
type GetAccountNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAccountNotFound creates GetAccountNotFound with default headers values
func NewGetAccountNotFound() *GetAccountNotFound {

	return &GetAccountNotFound{}
}

// WithPayload adds the payload to the get account not found response
func (o *GetAccountNotFound) WithPayload(payload *models.Error) *GetAccountNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account not found response
func (o *GetAccountNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAccountDefault Error

swagger:response getAccountDefault
*/
type GetAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAccountDefault creates GetAccountDefault with default headers values
func NewGetAccountDefault(code int) *GetAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get account default response
func (o *GetAccountDefault) WithStatusCode(code int) *GetAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get account default response
func (o *GetAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get account default response
func (o *GetAccountDefault) WithPayload(payload *models.Error) *GetAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account default response
func (o *GetAccountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
