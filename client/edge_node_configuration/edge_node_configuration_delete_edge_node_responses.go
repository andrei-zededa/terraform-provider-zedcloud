// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNodeConfigurationDeleteEdgeNodeReader is a Reader for the EdgeNodeConfigurationDeleteEdgeNode structure.
type EdgeNodeConfigurationDeleteEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationDeleteEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationDeleteEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationDeleteEdgeNodeOK creates a EdgeNodeConfigurationDeleteEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeOK() *EdgeNodeConfigurationDeleteEdgeNodeOK {
	return &EdgeNodeConfigurationDeleteEdgeNodeOK{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationDeleteEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration delete edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration delete edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration delete edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeUnauthorized creates a EdgeNodeConfigurationDeleteEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeUnauthorized() *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationDeleteEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationDeleteEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration delete edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration delete edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration delete edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeForbidden creates a EdgeNodeConfigurationDeleteEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeForbidden() *EdgeNodeConfigurationDeleteEdgeNodeForbidden {
	return &EdgeNodeConfigurationDeleteEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationDeleteEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration delete edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration delete edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration delete edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeNotFound creates a EdgeNodeConfigurationDeleteEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeNotFound() *EdgeNodeConfigurationDeleteEdgeNodeNotFound {
	return &EdgeNodeConfigurationDeleteEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationDeleteEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration delete edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration delete edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration delete edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeInternalServerError creates a EdgeNodeConfigurationDeleteEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeInternalServerError() *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationDeleteEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationDeleteEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration delete edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration delete edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration delete edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout() *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration delete edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration delete edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration delete edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration delete edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration delete edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] edgeNodeConfigurationDeleteEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeleteEdgeNodeDefault creates a EdgeNodeConfigurationDeleteEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationDeleteEdgeNodeDefault(code int) *EdgeNodeConfigurationDeleteEdgeNodeDefault {
	return &EdgeNodeConfigurationDeleteEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationDeleteEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationDeleteEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration delete edge node default response
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration delete edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration delete edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration delete edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration delete edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration delete edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] EdgeNodeConfiguration_DeleteEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] EdgeNodeConfiguration_DeleteEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeleteEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
