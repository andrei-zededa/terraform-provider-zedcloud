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

// EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader is a Reader for the EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstances structure.
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK struct {
	Payload *models.AppInstances
}

// IsSuccess returns true when this edge application instance configuration query edge application instances o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration query edge application instances o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration query edge application instances o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration query edge application instances o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration query edge application instances o k response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) GetPayload() *models.AppInstances {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration query edge application instances bad request response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration query edge application instances bad request response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances bad request response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration query edge application instances bad request response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration query edge application instances bad request response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge application instance configuration query edge application instances bad request response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) Code() int {
	return 400
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration query edge application instances unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration query edge application instances unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration query edge application instances unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration query edge application instances unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration query edge application instances unauthorized response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration query edge application instances forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration query edge application instances forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration query edge application instances forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration query edge application instances forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration query edge application instances forbidden response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration query edge application instances internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration query edge application instances internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration query edge application instances internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration query edge application instances internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration query edge application instances internal server error response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration query edge application instances gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration query edge application instances gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration query edge application instances gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration query edge application instances gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration query edge application instances gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration query edge application instances gateway timeout response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault(code int) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration query edge application instances default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration query edge application instances default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration query edge application instances default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration query edge application instances default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration query edge application instances default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration query edge application instances default response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] EdgeApplicationInstanceConfiguration_QueryEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] EdgeApplicationInstanceConfiguration_QueryEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
