// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceReader is a Reader for the EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstance structure.
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance o k response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance o k response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance o k response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration delete edge network instance o k response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration delete edge network instance o k response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network instance configuration delete edge network instance o k response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) Code() int {
	return 200
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance unauthorized response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance unauthorized response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance unauthorized response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration delete edge network instance unauthorized response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration delete edge network instance unauthorized response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network instance configuration delete edge network instance unauthorized response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance forbidden response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance forbidden response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance forbidden response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration delete edge network instance forbidden response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration delete edge network instance forbidden response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network instance configuration delete edge network instance forbidden response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance not found response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance not found response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance not found response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration delete edge network instance not found response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration delete edge network instance not found response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge network instance configuration delete edge network instance not found response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) Code() int {
	return 404
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance internal server error response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance internal server error response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance internal server error response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration delete edge network instance internal server error response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration delete edge network instance internal server error response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network instance configuration delete edge network instance internal server error response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout{}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance gateway timeout response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance gateway timeout response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration delete edge network instance gateway timeout response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration delete edge network instance gateway timeout response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration delete edge network instance gateway timeout response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network instance configuration delete edge network instance gateway timeout response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] edgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault creates a EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault with default headers values
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault(code int) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge network instance configuration delete edge network instance default response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network instance configuration delete edge network instance default response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network instance configuration delete edge network instance default response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network instance configuration delete edge network instance default response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network instance configuration delete edge network instance default response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network instance configuration delete edge network instance default response
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] EdgeNetworkInstanceConfiguration_DeleteEdgeNetworkInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/netinsts/id/{id}][%d] EdgeNetworkInstanceConfiguration_DeleteEdgeNetworkInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
