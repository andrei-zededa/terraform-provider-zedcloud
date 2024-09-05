// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// IdentityAccessManagementLoginReader is a Reader for the IdentityAccessManagementLogin structure.
type IdentityAccessManagementLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLoginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLoginGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLoginOK creates a IdentityAccessManagementLoginOK with default headers values
func NewIdentityAccessManagementLoginOK() *IdentityAccessManagementLoginOK {
	return &IdentityAccessManagementLoginOK{}
}

/*
IdentityAccessManagementLoginOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLoginOK struct {
	Payload *models.AAAFrontendLoginResponse
}

// IsSuccess returns true when this identity access management login o k response has a 2xx status code
func (o *IdentityAccessManagementLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management login o k response has a 3xx status code
func (o *IdentityAccessManagementLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login o k response has a 4xx status code
func (o *IdentityAccessManagementLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login o k response has a 5xx status code
func (o *IdentityAccessManagementLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login o k response a status code equal to that given
func (o *IdentityAccessManagementLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management login o k response
func (o *IdentityAccessManagementLoginOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementLoginOK) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginOK) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginOK) GetPayload() *models.AAAFrontendLoginResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AAAFrontendLoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginUnauthorized creates a IdentityAccessManagementLoginUnauthorized with default headers values
func NewIdentityAccessManagementLoginUnauthorized() *IdentityAccessManagementLoginUnauthorized {
	return &IdentityAccessManagementLoginUnauthorized{}
}

/*
IdentityAccessManagementLoginUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLoginUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login unauthorized response has a 2xx status code
func (o *IdentityAccessManagementLoginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login unauthorized response has a 3xx status code
func (o *IdentityAccessManagementLoginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login unauthorized response has a 4xx status code
func (o *IdentityAccessManagementLoginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login unauthorized response has a 5xx status code
func (o *IdentityAccessManagementLoginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login unauthorized response a status code equal to that given
func (o *IdentityAccessManagementLoginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management login unauthorized response
func (o *IdentityAccessManagementLoginUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForbidden creates a IdentityAccessManagementLoginForbidden with default headers values
func NewIdentityAccessManagementLoginForbidden() *IdentityAccessManagementLoginForbidden {
	return &IdentityAccessManagementLoginForbidden{}
}

/*
IdentityAccessManagementLoginForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLoginForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login forbidden response has a 2xx status code
func (o *IdentityAccessManagementLoginForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login forbidden response has a 3xx status code
func (o *IdentityAccessManagementLoginForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login forbidden response has a 4xx status code
func (o *IdentityAccessManagementLoginForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login forbidden response has a 5xx status code
func (o *IdentityAccessManagementLoginForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login forbidden response a status code equal to that given
func (o *IdentityAccessManagementLoginForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management login forbidden response
func (o *IdentityAccessManagementLoginForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementLoginForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginForbidden) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginInternalServerError creates a IdentityAccessManagementLoginInternalServerError with default headers values
func NewIdentityAccessManagementLoginInternalServerError() *IdentityAccessManagementLoginInternalServerError {
	return &IdentityAccessManagementLoginInternalServerError{}
}

/*
IdentityAccessManagementLoginInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLoginInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login internal server error response has a 2xx status code
func (o *IdentityAccessManagementLoginInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login internal server error response has a 3xx status code
func (o *IdentityAccessManagementLoginInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login internal server error response has a 4xx status code
func (o *IdentityAccessManagementLoginInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login internal server error response has a 5xx status code
func (o *IdentityAccessManagementLoginInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login internal server error response a status code equal to that given
func (o *IdentityAccessManagementLoginInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management login internal server error response
func (o *IdentityAccessManagementLoginInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementLoginInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginGatewayTimeout creates a IdentityAccessManagementLoginGatewayTimeout with default headers values
func NewIdentityAccessManagementLoginGatewayTimeout() *IdentityAccessManagementLoginGatewayTimeout {
	return &IdentityAccessManagementLoginGatewayTimeout{}
}

/*
IdentityAccessManagementLoginGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLoginGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementLoginGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementLoginGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementLoginGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementLoginGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementLoginGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management login gateway timeout response
func (o *IdentityAccessManagementLoginGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementLoginGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] identityAccessManagementLoginGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginDefault creates a IdentityAccessManagementLoginDefault with default headers values
func NewIdentityAccessManagementLoginDefault(code int) *IdentityAccessManagementLoginDefault {
	return &IdentityAccessManagementLoginDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementLoginDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLoginDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management login default response has a 2xx status code
func (o *IdentityAccessManagementLoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management login default response has a 3xx status code
func (o *IdentityAccessManagementLoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management login default response has a 4xx status code
func (o *IdentityAccessManagementLoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management login default response has a 5xx status code
func (o *IdentityAccessManagementLoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management login default response a status code equal to that given
func (o *IdentityAccessManagementLoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management login default response
func (o *IdentityAccessManagementLoginDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementLoginDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] IdentityAccessManagement_Login default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginDefault) String() string {
	return fmt.Sprintf("[POST /v1/login][%d] IdentityAccessManagement_Login default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
