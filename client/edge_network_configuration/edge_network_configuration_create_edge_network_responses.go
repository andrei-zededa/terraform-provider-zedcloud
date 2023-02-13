// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNetworkConfigurationCreateEdgeNetworkReader is a Reader for the EdgeNetworkConfigurationCreateEdgeNetwork structure.
type EdgeNetworkConfigurationCreateEdgeNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationCreateEdgeNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkOK creates a EdgeNetworkConfigurationCreateEdgeNetworkOK with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkOK() *EdgeNetworkConfigurationCreateEdgeNetworkOK {
	return &EdgeNetworkConfigurationCreateEdgeNetworkOK{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network o k response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network configuration create edge network o k response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network o k response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration create edge network o k response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration create edge network o k response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network configuration create edge network o k response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) Code() int {
	return 200
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest creates a EdgeNetworkConfigurationCreateEdgeNetworkBadRequest with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest() *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest {
	return &EdgeNetworkConfigurationCreateEdgeNetworkBadRequest{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network bad request response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network bad request response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network bad request response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration create edge network bad request response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration create edge network bad request response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge network configuration create edge network bad request response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) Code() int {
	return 400
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized creates a EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized() *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized {
	return &EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network unauthorized response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network unauthorized response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network unauthorized response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration create edge network unauthorized response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration create edge network unauthorized response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network configuration create edge network unauthorized response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden creates a EdgeNetworkConfigurationCreateEdgeNetworkForbidden with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden() *EdgeNetworkConfigurationCreateEdgeNetworkForbidden {
	return &EdgeNetworkConfigurationCreateEdgeNetworkForbidden{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network forbidden response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network forbidden response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network forbidden response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration create edge network forbidden response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration create edge network forbidden response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network configuration create edge network forbidden response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkConflict creates a EdgeNetworkConfigurationCreateEdgeNetworkConflict with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkConflict() *EdgeNetworkConfigurationCreateEdgeNetworkConflict {
	return &EdgeNetworkConfigurationCreateEdgeNetworkConflict{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network conflict response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network conflict response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network conflict response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration create edge network conflict response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration create edge network conflict response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge network configuration create edge network conflict response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) Code() int {
	return 409
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkConflict  %+v", 409, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkConflict  %+v", 409, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError creates a EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError() *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError {
	return &EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network internal server error response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network internal server error response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network internal server error response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration create edge network internal server error response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration create edge network internal server error response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network configuration create edge network internal server error response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout creates a EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout() *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout {
	return &EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout{}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration create edge network gateway timeout response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration create edge network gateway timeout response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration create edge network gateway timeout response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration create edge network gateway timeout response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration create edge network gateway timeout response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network configuration create edge network gateway timeout response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkDefault creates a EdgeNetworkConfigurationCreateEdgeNetworkDefault with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkDefault(code int) *EdgeNetworkConfigurationCreateEdgeNetworkDefault {
	return &EdgeNetworkConfigurationCreateEdgeNetworkDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge network configuration create edge network default response has a 2xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network configuration create edge network default response has a 3xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network configuration create edge network default response has a 4xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network configuration create edge network default response has a 5xx status code
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network configuration create edge network default response a status code equal to that given
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network configuration create edge network default response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] EdgeNetworkConfiguration_CreateEdgeNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) String() string {
	return fmt.Sprintf("[POST /v1/networks][%d] EdgeNetworkConfiguration_CreateEdgeNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
