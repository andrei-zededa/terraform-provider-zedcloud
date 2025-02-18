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

// GetListbyIdReader is a Reader for the ResourceGroupGetDeploymentListbyIdv2 structure.
type GetListbyIdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListbyIdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListbyIdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetListbyIdUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListbyIdForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListbyIdInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetListbyIdGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupGetDeploymentListbyIdv2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetListbyIdOK creates a GetListbyIdOK with default headers values
func NewGetListbyIdOK() *GetListbyIdOK {
	return &GetListbyIdOK{}
}

/*
GetListbyIdOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetListbyIdOK struct {
	Payload *models.Deployments
}

// IsSuccess returns true when this resource group get deployment listby idv2 o k response has a 2xx status code
func (o *GetListbyIdOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group get deployment listby idv2 o k response has a 3xx status code
func (o *GetListbyIdOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get deployment listby idv2 o k response has a 4xx status code
func (o *GetListbyIdOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get deployment listby idv2 o k response has a 5xx status code
func (o *GetListbyIdOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get deployment listby idv2 o k response a status code equal to that given
func (o *GetListbyIdOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource group get deployment listby idv2 o k response
func (o *GetListbyIdOK) Code() int {
	return 200
}

func (o *GetListbyIdOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdOK %s", 200, payload)
}

func (o *GetListbyIdOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdOK %s", 200, payload)
}

func (o *GetListbyIdOK) GetPayload() *models.Deployments {
	return o.Payload
}

func (o *GetListbyIdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListbyIdUnauthorized creates a GetListbyIdUnauthorized with default headers values
func NewGetListbyIdUnauthorized() *GetListbyIdUnauthorized {
	return &GetListbyIdUnauthorized{}
}

/*
GetListbyIdUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetListbyIdUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group get deployment listby idv2 unauthorized response has a 2xx status code
func (o *GetListbyIdUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get deployment listby idv2 unauthorized response has a 3xx status code
func (o *GetListbyIdUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get deployment listby idv2 unauthorized response has a 4xx status code
func (o *GetListbyIdUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group get deployment listby idv2 unauthorized response has a 5xx status code
func (o *GetListbyIdUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get deployment listby idv2 unauthorized response a status code equal to that given
func (o *GetListbyIdUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resource group get deployment listby idv2 unauthorized response
func (o *GetListbyIdUnauthorized) Code() int {
	return 401
}

func (o *GetListbyIdUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdUnauthorized %s", 401, payload)
}

func (o *GetListbyIdUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdUnauthorized %s", 401, payload)
}

func (o *GetListbyIdUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetListbyIdUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListbyIdForbidden creates a GetListbyIdForbidden with default headers values
func NewGetListbyIdForbidden() *GetListbyIdForbidden {
	return &GetListbyIdForbidden{}
}

/*
GetListbyIdForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetListbyIdForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group get deployment listby idv2 forbidden response has a 2xx status code
func (o *GetListbyIdForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get deployment listby idv2 forbidden response has a 3xx status code
func (o *GetListbyIdForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get deployment listby idv2 forbidden response has a 4xx status code
func (o *GetListbyIdForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group get deployment listby idv2 forbidden response has a 5xx status code
func (o *GetListbyIdForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get deployment listby idv2 forbidden response a status code equal to that given
func (o *GetListbyIdForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resource group get deployment listby idv2 forbidden response
func (o *GetListbyIdForbidden) Code() int {
	return 403
}

func (o *GetListbyIdForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdForbidden %s", 403, payload)
}

func (o *GetListbyIdForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdForbidden %s", 403, payload)
}

func (o *GetListbyIdForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetListbyIdForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListbyIdInternalServerError creates a GetListbyIdInternalServerError with default headers values
func NewGetListbyIdInternalServerError() *GetListbyIdInternalServerError {
	return &GetListbyIdInternalServerError{}
}

/*
GetListbyIdInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetListbyIdInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group get deployment listby idv2 internal server error response has a 2xx status code
func (o *GetListbyIdInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get deployment listby idv2 internal server error response has a 3xx status code
func (o *GetListbyIdInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get deployment listby idv2 internal server error response has a 4xx status code
func (o *GetListbyIdInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get deployment listby idv2 internal server error response has a 5xx status code
func (o *GetListbyIdInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group get deployment listby idv2 internal server error response a status code equal to that given
func (o *GetListbyIdInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource group get deployment listby idv2 internal server error response
func (o *GetListbyIdInternalServerError) Code() int {
	return 500
}

func (o *GetListbyIdInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdInternalServerError %s", 500, payload)
}

func (o *GetListbyIdInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdInternalServerError %s", 500, payload)
}

func (o *GetListbyIdInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetListbyIdInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListbyIdGatewayTimeout creates a GetListbyIdGatewayTimeout with default headers values
func NewGetListbyIdGatewayTimeout() *GetListbyIdGatewayTimeout {
	return &GetListbyIdGatewayTimeout{}
}

/*
GetListbyIdGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetListbyIdGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group get deployment listby idv2 gateway timeout response has a 2xx status code
func (o *GetListbyIdGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get deployment listby idv2 gateway timeout response has a 3xx status code
func (o *GetListbyIdGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get deployment listby idv2 gateway timeout response has a 4xx status code
func (o *GetListbyIdGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get deployment listby idv2 gateway timeout response has a 5xx status code
func (o *GetListbyIdGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group get deployment listby idv2 gateway timeout response a status code equal to that given
func (o *GetListbyIdGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the resource group get deployment listby idv2 gateway timeout response
func (o *GetListbyIdGatewayTimeout) Code() int {
	return 504
}

func (o *GetListbyIdGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdGatewayTimeout %s", 504, payload)
}

func (o *GetListbyIdGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListByIdGatewayTimeout %s", 504, payload)
}

func (o *GetListbyIdGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetListbyIdGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetDeploymentListbyIdv2Default creates a ResourceGroupGetDeploymentListbyIdv2Default with default headers values
func NewResourceGroupGetDeploymentListbyIdv2Default(code int) *GetListbyIdDefault {
	return &GetListbyIdDefault{
		_statusCode: code,
	}
}

/*
GetListbyIdDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetListbyIdDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this resource group get deployment listby idv2 default response has a 2xx status code
func (o *GetListbyIdDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group get deployment listby idv2 default response has a 3xx status code
func (o *GetListbyIdDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group get deployment listby idv2 default response has a 4xx status code
func (o *GetListbyIdDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group get deployment listby idv2 default response has a 5xx status code
func (o *GetListbyIdDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group get deployment listby idv2 default response a status code equal to that given
func (o *GetListbyIdDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource group get deployment listby idv2 default response
func (o *GetListbyIdDefault) Code() int {
	return o._statusCode
}

func (o *GetListbyIdDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListById default %s", o._statusCode, payload)
}

func (o *GetListbyIdDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/projects/id/{projectId}/deployments][%d] projectDeploymentGetListById default %s", o._statusCode, payload)
}

func (o *GetListbyIdDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetListbyIdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
