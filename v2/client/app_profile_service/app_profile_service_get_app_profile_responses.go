// Code generated by go-swagger; DO NOT EDIT.

package app_profile_service

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

// AppProfileServiceGetAppProfileReader is a Reader for the AppProfileServiceGetAppProfile structure.
type AppProfileServiceGetAppProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppProfileServiceGetAppProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppProfileServiceGetAppProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAppProfileServiceGetAppProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppProfileServiceGetAppProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppProfileServiceGetAppProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppProfileServiceGetAppProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewAppProfileServiceGetAppProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppProfileServiceGetAppProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppProfileServiceGetAppProfileOK creates a AppProfileServiceGetAppProfileOK with default headers values
func NewAppProfileServiceGetAppProfileOK() *AppProfileServiceGetAppProfileOK {
	return &AppProfileServiceGetAppProfileOK{}
}

/*
AppProfileServiceGetAppProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type AppProfileServiceGetAppProfileOK struct {
	Payload *models.AppProfileRead
}

// IsSuccess returns true when this app profile service get app profile o k response has a 2xx status code
func (o *AppProfileServiceGetAppProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app profile service get app profile o k response has a 3xx status code
func (o *AppProfileServiceGetAppProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile o k response has a 4xx status code
func (o *AppProfileServiceGetAppProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service get app profile o k response has a 5xx status code
func (o *AppProfileServiceGetAppProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service get app profile o k response a status code equal to that given
func (o *AppProfileServiceGetAppProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the app profile service get app profile o k response
func (o *AppProfileServiceGetAppProfileOK) Code() int {
	return 200
}

func (o *AppProfileServiceGetAppProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileOK %s", 200, payload)
}

func (o *AppProfileServiceGetAppProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileOK %s", 200, payload)
}

func (o *AppProfileServiceGetAppProfileOK) GetPayload() *models.AppProfileRead {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppProfileRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileUnauthorized creates a AppProfileServiceGetAppProfileUnauthorized with default headers values
func NewAppProfileServiceGetAppProfileUnauthorized() *AppProfileServiceGetAppProfileUnauthorized {
	return &AppProfileServiceGetAppProfileUnauthorized{}
}

/*
AppProfileServiceGetAppProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type AppProfileServiceGetAppProfileUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service get app profile unauthorized response has a 2xx status code
func (o *AppProfileServiceGetAppProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service get app profile unauthorized response has a 3xx status code
func (o *AppProfileServiceGetAppProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile unauthorized response has a 4xx status code
func (o *AppProfileServiceGetAppProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service get app profile unauthorized response has a 5xx status code
func (o *AppProfileServiceGetAppProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service get app profile unauthorized response a status code equal to that given
func (o *AppProfileServiceGetAppProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the app profile service get app profile unauthorized response
func (o *AppProfileServiceGetAppProfileUnauthorized) Code() int {
	return 401
}

func (o *AppProfileServiceGetAppProfileUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileUnauthorized %s", 401, payload)
}

func (o *AppProfileServiceGetAppProfileUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileUnauthorized %s", 401, payload)
}

func (o *AppProfileServiceGetAppProfileUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileForbidden creates a AppProfileServiceGetAppProfileForbidden with default headers values
func NewAppProfileServiceGetAppProfileForbidden() *AppProfileServiceGetAppProfileForbidden {
	return &AppProfileServiceGetAppProfileForbidden{}
}

/*
AppProfileServiceGetAppProfileForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type AppProfileServiceGetAppProfileForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service get app profile forbidden response has a 2xx status code
func (o *AppProfileServiceGetAppProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service get app profile forbidden response has a 3xx status code
func (o *AppProfileServiceGetAppProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile forbidden response has a 4xx status code
func (o *AppProfileServiceGetAppProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service get app profile forbidden response has a 5xx status code
func (o *AppProfileServiceGetAppProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service get app profile forbidden response a status code equal to that given
func (o *AppProfileServiceGetAppProfileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the app profile service get app profile forbidden response
func (o *AppProfileServiceGetAppProfileForbidden) Code() int {
	return 403
}

func (o *AppProfileServiceGetAppProfileForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileForbidden %s", 403, payload)
}

func (o *AppProfileServiceGetAppProfileForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileForbidden %s", 403, payload)
}

func (o *AppProfileServiceGetAppProfileForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileNotFound creates a AppProfileServiceGetAppProfileNotFound with default headers values
func NewAppProfileServiceGetAppProfileNotFound() *AppProfileServiceGetAppProfileNotFound {
	return &AppProfileServiceGetAppProfileNotFound{}
}

/*
AppProfileServiceGetAppProfileNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type AppProfileServiceGetAppProfileNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service get app profile not found response has a 2xx status code
func (o *AppProfileServiceGetAppProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service get app profile not found response has a 3xx status code
func (o *AppProfileServiceGetAppProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile not found response has a 4xx status code
func (o *AppProfileServiceGetAppProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service get app profile not found response has a 5xx status code
func (o *AppProfileServiceGetAppProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service get app profile not found response a status code equal to that given
func (o *AppProfileServiceGetAppProfileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app profile service get app profile not found response
func (o *AppProfileServiceGetAppProfileNotFound) Code() int {
	return 404
}

func (o *AppProfileServiceGetAppProfileNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileNotFound %s", 404, payload)
}

func (o *AppProfileServiceGetAppProfileNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileNotFound %s", 404, payload)
}

func (o *AppProfileServiceGetAppProfileNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileInternalServerError creates a AppProfileServiceGetAppProfileInternalServerError with default headers values
func NewAppProfileServiceGetAppProfileInternalServerError() *AppProfileServiceGetAppProfileInternalServerError {
	return &AppProfileServiceGetAppProfileInternalServerError{}
}

/*
AppProfileServiceGetAppProfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type AppProfileServiceGetAppProfileInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service get app profile internal server error response has a 2xx status code
func (o *AppProfileServiceGetAppProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service get app profile internal server error response has a 3xx status code
func (o *AppProfileServiceGetAppProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile internal server error response has a 4xx status code
func (o *AppProfileServiceGetAppProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service get app profile internal server error response has a 5xx status code
func (o *AppProfileServiceGetAppProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app profile service get app profile internal server error response a status code equal to that given
func (o *AppProfileServiceGetAppProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the app profile service get app profile internal server error response
func (o *AppProfileServiceGetAppProfileInternalServerError) Code() int {
	return 500
}

func (o *AppProfileServiceGetAppProfileInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileInternalServerError %s", 500, payload)
}

func (o *AppProfileServiceGetAppProfileInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileInternalServerError %s", 500, payload)
}

func (o *AppProfileServiceGetAppProfileInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileGatewayTimeout creates a AppProfileServiceGetAppProfileGatewayTimeout with default headers values
func NewAppProfileServiceGetAppProfileGatewayTimeout() *AppProfileServiceGetAppProfileGatewayTimeout {
	return &AppProfileServiceGetAppProfileGatewayTimeout{}
}

/*
AppProfileServiceGetAppProfileGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type AppProfileServiceGetAppProfileGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service get app profile gateway timeout response has a 2xx status code
func (o *AppProfileServiceGetAppProfileGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service get app profile gateway timeout response has a 3xx status code
func (o *AppProfileServiceGetAppProfileGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service get app profile gateway timeout response has a 4xx status code
func (o *AppProfileServiceGetAppProfileGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service get app profile gateway timeout response has a 5xx status code
func (o *AppProfileServiceGetAppProfileGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this app profile service get app profile gateway timeout response a status code equal to that given
func (o *AppProfileServiceGetAppProfileGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the app profile service get app profile gateway timeout response
func (o *AppProfileServiceGetAppProfileGatewayTimeout) Code() int {
	return 504
}

func (o *AppProfileServiceGetAppProfileGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileGatewayTimeout %s", 504, payload)
}

func (o *AppProfileServiceGetAppProfileGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] appProfileServiceGetAppProfileGatewayTimeout %s", 504, payload)
}

func (o *AppProfileServiceGetAppProfileGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceGetAppProfileDefault creates a AppProfileServiceGetAppProfileDefault with default headers values
func NewAppProfileServiceGetAppProfileDefault(code int) *AppProfileServiceGetAppProfileDefault {
	return &AppProfileServiceGetAppProfileDefault{
		_statusCode: code,
	}
}

/*
AppProfileServiceGetAppProfileDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AppProfileServiceGetAppProfileDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this app profile service get app profile default response has a 2xx status code
func (o *AppProfileServiceGetAppProfileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this app profile service get app profile default response has a 3xx status code
func (o *AppProfileServiceGetAppProfileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this app profile service get app profile default response has a 4xx status code
func (o *AppProfileServiceGetAppProfileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this app profile service get app profile default response has a 5xx status code
func (o *AppProfileServiceGetAppProfileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this app profile service get app profile default response a status code equal to that given
func (o *AppProfileServiceGetAppProfileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the app profile service get app profile default response
func (o *AppProfileServiceGetAppProfileDefault) Code() int {
	return o._statusCode
}

func (o *AppProfileServiceGetAppProfileDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] AppProfileService_GetAppProfile default %s", o._statusCode, payload)
}

func (o *AppProfileServiceGetAppProfileDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/appprofiles/id/{id}][%d] AppProfileService_GetAppProfile default %s", o._statusCode, payload)
}

func (o *AppProfileServiceGetAppProfileDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AppProfileServiceGetAppProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
