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

// EdgeNodeStatusGetEdgeNodeRawStatusByNameReader is a Reader for the EdgeNodeStatusGetEdgeNodeRawStatusByName structure.
type EdgeNodeStatusGetEdgeNodeRawStatusByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeRawStatusByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameOK creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameOK() *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameOK{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameOK struct {
	Payload *models.DeviceRawMetrics
}

// IsSuccess returns true when this edge node status get edge node raw status by name o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node raw status by name o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node raw status by name o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node raw status by name o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) GetPayload() *models.DeviceRawMetrics {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRawMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized() *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node raw status by name unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node raw status by name unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node raw status by name unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node raw status by name unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden() *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node raw status by name forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node raw status by name forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node raw status by name forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node raw status by name forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound() *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node raw status by name not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node raw status by name not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node raw status by name not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node raw status by name not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError() *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node raw status by name internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node raw status by name internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node raw status by name internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node raw status by name internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout() *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node raw status by name gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node raw status by name gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node raw status by name gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node raw status by name gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node raw status by name gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] edgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameDefault creates a EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameDefault(code int) *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node raw status by name default response
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node raw status by name default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node raw status by name default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node raw status by name default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node raw status by name default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node raw status by name default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] EdgeNodeStatus_GetEdgeNodeRawStatusByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/metrics/raw][%d] EdgeNodeStatus_GetEdgeNodeRawStatusByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
