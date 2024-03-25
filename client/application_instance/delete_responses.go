package application_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceReader is a Reader for the EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstance structure.
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration delete edge application instance o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration delete edge application instance o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration delete edge application instance o k response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration delete edge application instance unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration delete edge application instance unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration delete edge application instance unauthorized response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration delete edge application instance forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration delete edge application instance forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration delete edge application instance forbidden response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance not found response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance not found response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance not found response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration delete edge application instance not found response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration delete edge application instance not found response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application instance configuration delete edge application instance not found response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration delete edge application instance internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration delete edge application instance internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration delete edge application instance internal server error response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout() *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration delete edge application instance gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration delete edge application instance gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration delete edge application instance gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration delete edge application instance gateway timeout response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault creates a EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault with default headers values
func NewEdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault(code int) *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault {
	return &EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration delete edge application instance default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration delete edge application instance default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration delete edge application instance default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration delete edge application instance default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration delete edge application instance default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration delete edge application instance default response
func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] EdgeApplicationInstanceConfiguration_DeleteEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] EdgeApplicationInstanceConfiguration_DeleteEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeleteEdgeApplicationInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
