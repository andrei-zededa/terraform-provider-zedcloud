// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNodeStatusGetEdgeNodeStatusReader is a Reader for the EdgeNodeStatusGetEdgeNodeStatus structure.
type EdgeNodeStatusGetEdgeNodeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeStatusOK creates a EdgeNodeStatusGetEdgeNodeStatusOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusOK() *EdgeNodeStatusGetEdgeNodeStatusOK {
	return &EdgeNodeStatusGetEdgeNodeStatusOK{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeStatusOK struct {
	Payload *models.DeviceStatus
}

// IsSuccess returns true when this edge node status get edge node status o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node status o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node status o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node status o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusOK) GetPayload() *models.DeviceStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusUnauthorized creates a EdgeNodeStatusGetEdgeNodeStatusUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusUnauthorized() *EdgeNodeStatusGetEdgeNodeStatusUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeStatusUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeStatusUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node status unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node status unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node status unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node status unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusForbidden creates a EdgeNodeStatusGetEdgeNodeStatusForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusForbidden() *EdgeNodeStatusGetEdgeNodeStatusForbidden {
	return &EdgeNodeStatusGetEdgeNodeStatusForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeStatusForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node status forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node status forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node status forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node status forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusNotFound creates a EdgeNodeStatusGetEdgeNodeStatusNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusNotFound() *EdgeNodeStatusGetEdgeNodeStatusNotFound {
	return &EdgeNodeStatusGetEdgeNodeStatusNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeStatusNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node status not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node status not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node status not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node status not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusInternalServerError creates a EdgeNodeStatusGetEdgeNodeStatusInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusInternalServerError() *EdgeNodeStatusGetEdgeNodeStatusInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeStatusInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeStatusInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node status internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node status internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node status internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node status internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusGatewayTimeout() *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node status gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node status gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node status gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node status gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node status gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] edgeNodeStatusGetEdgeNodeStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeStatusDefault creates a EdgeNodeStatusGetEdgeNodeStatusDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeStatusDefault(code int) *EdgeNodeStatusGetEdgeNodeStatusDefault {
	return &EdgeNodeStatusGetEdgeNodeStatusDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeStatusDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node status default response
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node status default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node status default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node status default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node status default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node status default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] EdgeNodeStatus_GetEdgeNodeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status][%d] EdgeNodeStatus_GetEdgeNodeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
