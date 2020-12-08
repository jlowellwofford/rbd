// Code generated by go-swagger; DO NOT EDIT.

package rbds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bensallen/rbd/models"
)

// GetRbdOKCode is the HTTP code returned for type GetRbdOK
const GetRbdOKCode int = 200

/*GetRbdOK RBD entry

swagger:response getRbdOK
*/
type GetRbdOK struct {

	/*
	  In: Body
	*/
	Payload *models.Rbd `json:"body,omitempty"`
}

// NewGetRbdOK creates GetRbdOK with default headers values
func NewGetRbdOK() *GetRbdOK {

	return &GetRbdOK{}
}

// WithPayload adds the payload to the get rbd o k response
func (o *GetRbdOK) WithPayload(payload *models.Rbd) *GetRbdOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rbd o k response
func (o *GetRbdOK) SetPayload(payload *models.Rbd) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRbdOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRbdDefault error

swagger:response getRbdDefault
*/
type GetRbdDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRbdDefault creates GetRbdDefault with default headers values
func NewGetRbdDefault(code int) *GetRbdDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRbdDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get rbd default response
func (o *GetRbdDefault) WithStatusCode(code int) *GetRbdDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get rbd default response
func (o *GetRbdDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get rbd default response
func (o *GetRbdDefault) WithPayload(payload *models.Error) *GetRbdDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rbd default response
func (o *GetRbdDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRbdDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}