// Code generated by go-swagger; DO NOT EDIT.

package scalesets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
	garm_params "github.com/cloudbase/garm/params"
)

// GetScaleSetReader is a Reader for the GetScaleSet structure.
type GetScaleSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScaleSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScaleSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetScaleSetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScaleSetOK creates a GetScaleSetOK with default headers values
func NewGetScaleSetOK() *GetScaleSetOK {
	return &GetScaleSetOK{}
}

/*
GetScaleSetOK describes a response with status code 200, with default header values.

ScaleSet
*/
type GetScaleSetOK struct {
	Payload garm_params.ScaleSet
}

// IsSuccess returns true when this get scale set o k response has a 2xx status code
func (o *GetScaleSetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scale set o k response has a 3xx status code
func (o *GetScaleSetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scale set o k response has a 4xx status code
func (o *GetScaleSetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scale set o k response has a 5xx status code
func (o *GetScaleSetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scale set o k response a status code equal to that given
func (o *GetScaleSetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scale set o k response
func (o *GetScaleSetOK) Code() int {
	return 200
}

func (o *GetScaleSetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /scalesets/{scalesetID}][%d] getScaleSetOK %s", 200, payload)
}

func (o *GetScaleSetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /scalesets/{scalesetID}][%d] getScaleSetOK %s", 200, payload)
}

func (o *GetScaleSetOK) GetPayload() garm_params.ScaleSet {
	return o.Payload
}

func (o *GetScaleSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScaleSetDefault creates a GetScaleSetDefault with default headers values
func NewGetScaleSetDefault(code int) *GetScaleSetDefault {
	return &GetScaleSetDefault{
		_statusCode: code,
	}
}

/*
GetScaleSetDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type GetScaleSetDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this get scale set default response has a 2xx status code
func (o *GetScaleSetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get scale set default response has a 3xx status code
func (o *GetScaleSetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get scale set default response has a 4xx status code
func (o *GetScaleSetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get scale set default response has a 5xx status code
func (o *GetScaleSetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get scale set default response a status code equal to that given
func (o *GetScaleSetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get scale set default response
func (o *GetScaleSetDefault) Code() int {
	return o._statusCode
}

func (o *GetScaleSetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /scalesets/{scalesetID}][%d] GetScaleSet default %s", o._statusCode, payload)
}

func (o *GetScaleSetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /scalesets/{scalesetID}][%d] GetScaleSet default %s", o._statusCode, payload)
}

func (o *GetScaleSetDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *GetScaleSetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
