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

// EdgeNodeConfigurationRebootReader is a Reader for the EdgeNodeConfigurationReboot structure.
type EdgeNodeConfigurationRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationRebootConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationRebootGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationRebootDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationRebootOK creates a EdgeNodeConfigurationRebootOK with default headers values
func NewEdgeNodeConfigurationRebootOK() *EdgeNodeConfigurationRebootOK {
	return &EdgeNodeConfigurationRebootOK{}
}

/*
EdgeNodeConfigurationRebootOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationRebootOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot o k response has a 2xx status code
func (o *EdgeNodeConfigurationRebootOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration reboot o k response has a 3xx status code
func (o *EdgeNodeConfigurationRebootOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot o k response has a 4xx status code
func (o *EdgeNodeConfigurationRebootOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration reboot o k response has a 5xx status code
func (o *EdgeNodeConfigurationRebootOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration reboot o k response a status code equal to that given
func (o *EdgeNodeConfigurationRebootOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationRebootOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationRebootOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationRebootOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootUnauthorized creates a EdgeNodeConfigurationRebootUnauthorized with default headers values
func NewEdgeNodeConfigurationRebootUnauthorized() *EdgeNodeConfigurationRebootUnauthorized {
	return &EdgeNodeConfigurationRebootUnauthorized{}
}

/*
EdgeNodeConfigurationRebootUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationRebootUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationRebootUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationRebootUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationRebootUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration reboot unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationRebootUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration reboot unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationRebootUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationRebootUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationRebootUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationRebootUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootForbidden creates a EdgeNodeConfigurationRebootForbidden with default headers values
func NewEdgeNodeConfigurationRebootForbidden() *EdgeNodeConfigurationRebootForbidden {
	return &EdgeNodeConfigurationRebootForbidden{}
}

/*
EdgeNodeConfigurationRebootForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationRebootForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationRebootForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationRebootForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationRebootForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration reboot forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationRebootForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration reboot forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationRebootForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationRebootForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationRebootForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationRebootForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootNotFound creates a EdgeNodeConfigurationRebootNotFound with default headers values
func NewEdgeNodeConfigurationRebootNotFound() *EdgeNodeConfigurationRebootNotFound {
	return &EdgeNodeConfigurationRebootNotFound{}
}

/*
EdgeNodeConfigurationRebootNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationRebootNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot not found response has a 2xx status code
func (o *EdgeNodeConfigurationRebootNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot not found response has a 3xx status code
func (o *EdgeNodeConfigurationRebootNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot not found response has a 4xx status code
func (o *EdgeNodeConfigurationRebootNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration reboot not found response has a 5xx status code
func (o *EdgeNodeConfigurationRebootNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration reboot not found response a status code equal to that given
func (o *EdgeNodeConfigurationRebootNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationRebootNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationRebootNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationRebootNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootConflict creates a EdgeNodeConfigurationRebootConflict with default headers values
func NewEdgeNodeConfigurationRebootConflict() *EdgeNodeConfigurationRebootConflict {
	return &EdgeNodeConfigurationRebootConflict{}
}

/*
EdgeNodeConfigurationRebootConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationRebootConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot conflict response has a 2xx status code
func (o *EdgeNodeConfigurationRebootConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot conflict response has a 3xx status code
func (o *EdgeNodeConfigurationRebootConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot conflict response has a 4xx status code
func (o *EdgeNodeConfigurationRebootConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration reboot conflict response has a 5xx status code
func (o *EdgeNodeConfigurationRebootConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration reboot conflict response a status code equal to that given
func (o *EdgeNodeConfigurationRebootConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationRebootConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationRebootConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationRebootConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootInternalServerError creates a EdgeNodeConfigurationRebootInternalServerError with default headers values
func NewEdgeNodeConfigurationRebootInternalServerError() *EdgeNodeConfigurationRebootInternalServerError {
	return &EdgeNodeConfigurationRebootInternalServerError{}
}

/*
EdgeNodeConfigurationRebootInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationRebootInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationRebootInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationRebootInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationRebootInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration reboot internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationRebootInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration reboot internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationRebootInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationRebootInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationRebootInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationRebootInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootGatewayTimeout creates a EdgeNodeConfigurationRebootGatewayTimeout with default headers values
func NewEdgeNodeConfigurationRebootGatewayTimeout() *EdgeNodeConfigurationRebootGatewayTimeout {
	return &EdgeNodeConfigurationRebootGatewayTimeout{}
}

/*
EdgeNodeConfigurationRebootGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationRebootGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration reboot gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationRebootGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration reboot gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationRebootGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration reboot gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationRebootGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration reboot gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationRebootGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration reboot gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationRebootGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootDefault creates a EdgeNodeConfigurationRebootDefault with default headers values
func NewEdgeNodeConfigurationRebootDefault(code int) *EdgeNodeConfigurationRebootDefault {
	return &EdgeNodeConfigurationRebootDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationRebootDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationRebootDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration reboot default response
func (o *EdgeNodeConfigurationRebootDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration reboot default response has a 2xx status code
func (o *EdgeNodeConfigurationRebootDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration reboot default response has a 3xx status code
func (o *EdgeNodeConfigurationRebootDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration reboot default response has a 4xx status code
func (o *EdgeNodeConfigurationRebootDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration reboot default response has a 5xx status code
func (o *EdgeNodeConfigurationRebootDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration reboot default response a status code equal to that given
func (o *EdgeNodeConfigurationRebootDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationRebootDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] EdgeNodeConfiguration_Reboot default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationRebootDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] EdgeNodeConfiguration_Reboot default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationRebootDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
