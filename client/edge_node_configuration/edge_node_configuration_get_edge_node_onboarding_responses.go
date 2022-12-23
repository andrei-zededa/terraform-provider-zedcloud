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

// EdgeNodeConfigurationGetEdgeNodeOnboardingReader is a Reader for the EdgeNodeConfigurationGetEdgeNodeOnboarding structure.
type EdgeNodeConfigurationGetEdgeNodeOnboardingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationGetEdgeNodeOnboardingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingOK creates a EdgeNodeConfigurationGetEdgeNodeOnboardingOK with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingOK() *EdgeNodeConfigurationGetEdgeNodeOnboardingOK {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingOK{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingOK struct {
	Payload *models.DeviceConfig
}

// IsSuccess returns true when this edge node configuration get edge node onboarding o k response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration get edge node onboarding o k response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding o k response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node onboarding o k response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node onboarding o k response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) GetPayload() *models.DeviceConfig {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized creates a EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized() *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node onboarding unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node onboarding unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node onboarding unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node onboarding unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingForbidden creates a EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingForbidden() *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node onboarding forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node onboarding forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node onboarding forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node onboarding forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingNotFound creates a EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingNotFound() *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node onboarding not found response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node onboarding not found response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding not found response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node onboarding not found response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node onboarding not found response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError creates a EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError() *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node onboarding internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node onboarding internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node onboarding internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node onboarding internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout creates a EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout() *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node onboarding gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node onboarding gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node onboarding gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node onboarding gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node onboarding gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] edgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingDefault creates a EdgeNodeConfigurationGetEdgeNodeOnboardingDefault with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingDefault(code int) *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration get edge node onboarding default response
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration get edge node onboarding default response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration get edge node onboarding default response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration get edge node onboarding default response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration get edge node onboarding default response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration get edge node onboarding default response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] EdgeNodeConfiguration_GetEdgeNodeOnboarding default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] EdgeNodeConfiguration_GetEdgeNodeOnboarding default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
