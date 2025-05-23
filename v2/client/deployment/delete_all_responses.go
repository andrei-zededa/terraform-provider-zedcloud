// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// DeleteAllReader is a Reader for the DeleteAll structure.
type DeleteAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAllGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAllOK creates a DeleteAllOK with default headers values
func NewDeleteAllOK() *DeleteAllOK {
	return &DeleteAllOK{}
}

/*
DeleteAllOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAllOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group all v2 o k response has a 2xx status code
func (o *DeleteAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group delete resource group all v2 o k response has a 3xx status code
func (o *DeleteAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group all v2 o k response has a 4xx status code
func (o *DeleteAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group all v2 o k response has a 5xx status code
func (o *DeleteAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group all v2 o k response a status code equal to that given
func (o *DeleteAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource group delete resource group all v2 o k response
func (o *DeleteAllOK) Code() int {
	return 200
}

func (o *DeleteAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllOK %s", 200, payload)
}

func (o *DeleteAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllOK %s", 200, payload)
}

func (o *DeleteAllOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllUnauthorized creates a DeleteAllUnauthorized with default headers values
func NewDeleteAllUnauthorized() *DeleteAllUnauthorized {
	return &DeleteAllUnauthorized{}
}

/*
DeleteAllUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteAllUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group all v2 unauthorized response has a 2xx status code
func (o *DeleteAllUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group all v2 unauthorized response has a 3xx status code
func (o *DeleteAllUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group all v2 unauthorized response has a 4xx status code
func (o *DeleteAllUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group delete resource group all v2 unauthorized response has a 5xx status code
func (o *DeleteAllUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group all v2 unauthorized response a status code equal to that given
func (o *DeleteAllUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resource group delete resource group all v2 unauthorized response
func (o *DeleteAllUnauthorized) Code() int {
	return 401
}

func (o *DeleteAllUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllUnauthorized %s", 401, payload)
}

func (o *DeleteAllUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllUnauthorized %s", 401, payload)
}

func (o *DeleteAllUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllForbidden creates a DeleteAllForbidden with default headers values
func NewDeleteAllForbidden() *DeleteAllForbidden {
	return &DeleteAllForbidden{}
}

/*
DeleteAllForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeleteAllForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group all v2 forbidden response has a 2xx status code
func (o *DeleteAllForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group all v2 forbidden response has a 3xx status code
func (o *DeleteAllForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group all v2 forbidden response has a 4xx status code
func (o *DeleteAllForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group delete resource group all v2 forbidden response has a 5xx status code
func (o *DeleteAllForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group all v2 forbidden response a status code equal to that given
func (o *DeleteAllForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resource group delete resource group all v2 forbidden response
func (o *DeleteAllForbidden) Code() int {
	return 403
}

func (o *DeleteAllForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllForbidden %s", 403, payload)
}

func (o *DeleteAllForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllForbidden %s", 403, payload)
}

func (o *DeleteAllForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllInternalServerError creates a DeleteAllInternalServerError with default headers values
func NewDeleteAllInternalServerError() *DeleteAllInternalServerError {
	return &DeleteAllInternalServerError{}
}

/*
DeleteAllInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteAllInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group all v2 internal server error response has a 2xx status code
func (o *DeleteAllInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group all v2 internal server error response has a 3xx status code
func (o *DeleteAllInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group all v2 internal server error response has a 4xx status code
func (o *DeleteAllInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group all v2 internal server error response has a 5xx status code
func (o *DeleteAllInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group delete resource group all v2 internal server error response a status code equal to that given
func (o *DeleteAllInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource group delete resource group all v2 internal server error response
func (o *DeleteAllInternalServerError) Code() int {
	return 500
}

func (o *DeleteAllInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllInternalServerError %s", 500, payload)
}

func (o *DeleteAllInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllInternalServerError %s", 500, payload)
}

func (o *DeleteAllInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllGatewayTimeout creates a DeleteAllGatewayTimeout with default headers values
func NewDeleteAllGatewayTimeout() *DeleteAllGatewayTimeout {
	return &DeleteAllGatewayTimeout{}
}

/*
DeleteAllGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteAllGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group all v2 gateway timeout response has a 2xx status code
func (o *DeleteAllGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group all v2 gateway timeout response has a 3xx status code
func (o *DeleteAllGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group all v2 gateway timeout response has a 4xx status code
func (o *DeleteAllGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group all v2 gateway timeout response has a 5xx status code
func (o *DeleteAllGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group delete resource group all v2 gateway timeout response a status code equal to that given
func (o *DeleteAllGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the resource group delete resource group all v2 gateway timeout response
func (o *DeleteAllGatewayTimeout) Code() int {
	return 504
}

func (o *DeleteAllGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllGatewayTimeout %s", 504, payload)
}

func (o *DeleteAllGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAllGatewayTimeout %s", 504, payload)
}

func (o *DeleteAllGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteAllGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllDefault creates a DeleteAllDefault with default headers values
func NewDeleteAllDefault(code int) *DeleteAllDefault {
	return &DeleteAllDefault{
		_statusCode: code,
	}
}

/*
DeleteAllDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAllDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this resource group delete resource group all v2 default response has a 2xx status code
func (o *DeleteAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group delete resource group all v2 default response has a 3xx status code
func (o *DeleteAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group delete resource group all v2 default response has a 4xx status code
func (o *DeleteAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group delete resource group all v2 default response has a 5xx status code
func (o *DeleteAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group delete resource group all v2 default response a status code equal to that given
func (o *DeleteAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource group delete resource group all v2 default response
func (o *DeleteAllDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAll default %s", o._statusCode, payload)
}

func (o *DeleteAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments][%d] projectDeploymentDeleteAll default %s", o._statusCode, payload)
}

func (o *DeleteAllDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
