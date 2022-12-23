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

// EdgeNodeStatusGetEdgeNodeResourceMetricsByIDReader is a Reader for the EdgeNodeStatusGetEdgeNodeResourceMetricsByID structure.
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK struct {
	Payload *models.MetricQueryResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node resource metrics by Id o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) GetPayload() *models.MetricQueryResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node resource metrics by Id unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node resource metrics by Id forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node resource metrics by Id not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node resource metrics by Id internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout() *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node resource metrics by Id gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] edgeNodeStatusGetEdgeNodeResourceMetricsByIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault creates a EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault(code int) *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault {
	return &EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node resource metrics by Id default response
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node resource metrics by Id default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node resource metrics by Id default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node resource metrics by Id default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node resource metrics by Id default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node resource metrics by Id default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] EdgeNodeStatus_GetEdgeNodeResourceMetricsById default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/timeSeries/{mType}][%d] EdgeNodeStatus_GetEdgeNodeResourceMetricsById default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeResourceMetricsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
