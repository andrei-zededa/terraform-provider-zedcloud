// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// DatastoreConfigurationGetDatastoreReader is a Reader for the DatastoreConfigurationGetDatastore structure.
type DatastoreConfigurationGetDatastoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatastoreConfigurationGetDatastoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatastoreConfigurationGetDatastoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDatastoreConfigurationGetDatastoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatastoreConfigurationGetDatastoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatastoreConfigurationGetDatastoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatastoreConfigurationGetDatastoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDatastoreConfigurationGetDatastoreGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDatastoreConfigurationGetDatastoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDatastoreConfigurationGetDatastoreOK creates a DatastoreConfigurationGetDatastoreOK with default headers values
func NewDatastoreConfigurationGetDatastoreOK() *DatastoreConfigurationGetDatastoreOK {
	return &DatastoreConfigurationGetDatastoreOK{}
}

/*
DatastoreConfigurationGetDatastoreOK describes a response with status code 200, with default header values.

A successful response.
*/
type DatastoreConfigurationGetDatastoreOK struct {
	Payload *models.Datastore
}

// IsSuccess returns true when this datastore configuration get datastore o k response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datastore configuration get datastore o k response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore o k response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore o k response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore o k response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datastore configuration get datastore o k response
func (o *DatastoreConfigurationGetDatastoreOK) Code() int {
	return 200
}

func (o *DatastoreConfigurationGetDatastoreOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreOK) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreOK) GetPayload() *models.Datastore {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datastore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreUnauthorized creates a DatastoreConfigurationGetDatastoreUnauthorized with default headers values
func NewDatastoreConfigurationGetDatastoreUnauthorized() *DatastoreConfigurationGetDatastoreUnauthorized {
	return &DatastoreConfigurationGetDatastoreUnauthorized{}
}

/*
DatastoreConfigurationGetDatastoreUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DatastoreConfigurationGetDatastoreUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore unauthorized response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore unauthorized response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore unauthorized response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore unauthorized response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore unauthorized response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datastore configuration get datastore unauthorized response
func (o *DatastoreConfigurationGetDatastoreUnauthorized) Code() int {
	return 401
}

func (o *DatastoreConfigurationGetDatastoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreForbidden creates a DatastoreConfigurationGetDatastoreForbidden with default headers values
func NewDatastoreConfigurationGetDatastoreForbidden() *DatastoreConfigurationGetDatastoreForbidden {
	return &DatastoreConfigurationGetDatastoreForbidden{}
}

/*
DatastoreConfigurationGetDatastoreForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DatastoreConfigurationGetDatastoreForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore forbidden response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore forbidden response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore forbidden response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore forbidden response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore forbidden response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datastore configuration get datastore forbidden response
func (o *DatastoreConfigurationGetDatastoreForbidden) Code() int {
	return 403
}

func (o *DatastoreConfigurationGetDatastoreForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreForbidden) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreNotFound creates a DatastoreConfigurationGetDatastoreNotFound with default headers values
func NewDatastoreConfigurationGetDatastoreNotFound() *GetByIDNotFound {
	return &GetByIDNotFound{}
}

/*
GetByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetByIDNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore not found response has a 2xx status code
func (o *GetByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore not found response has a 3xx status code
func (o *GetByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore not found response has a 4xx status code
func (o *GetByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore not found response has a 5xx status code
func (o *GetByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore not found response a status code equal to that given
func (o *GetByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the datastore configuration get datastore not found response
func (o *GetByIDNotFound) Code() int {
	return 404
}

func (o *GetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreNotFound  %+v", 404, o.Payload)
}

func (o *GetByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreNotFound  %+v", 404, o.Payload)
}

func (o *GetByIDNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreInternalServerError creates a DatastoreConfigurationGetDatastoreInternalServerError with default headers values
func NewDatastoreConfigurationGetDatastoreInternalServerError() *DatastoreConfigurationGetDatastoreInternalServerError {
	return &DatastoreConfigurationGetDatastoreInternalServerError{}
}

/*
DatastoreConfigurationGetDatastoreInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DatastoreConfigurationGetDatastoreInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore internal server error response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore internal server error response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore internal server error response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore internal server error response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration get datastore internal server error response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datastore configuration get datastore internal server error response
func (o *DatastoreConfigurationGetDatastoreInternalServerError) Code() int {
	return 500
}

func (o *DatastoreConfigurationGetDatastoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreGatewayTimeout creates a DatastoreConfigurationGetDatastoreGatewayTimeout with default headers values
func NewDatastoreConfigurationGetDatastoreGatewayTimeout() *DatastoreConfigurationGetDatastoreGatewayTimeout {
	return &DatastoreConfigurationGetDatastoreGatewayTimeout{}
}

/*
DatastoreConfigurationGetDatastoreGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DatastoreConfigurationGetDatastoreGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore gateway timeout response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore gateway timeout response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore gateway timeout response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore gateway timeout response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration get datastore gateway timeout response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the datastore configuration get datastore gateway timeout response
func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) Code() int {
	return 504
}

func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] datastoreConfigurationGetDatastoreGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreDefault creates a DatastoreConfigurationGetDatastoreDefault with default headers values
func NewDatastoreConfigurationGetDatastoreDefault(code int) *DatastoreConfigurationGetDatastoreDefault {
	return &DatastoreConfigurationGetDatastoreDefault{
		_statusCode: code,
	}
}

/*
DatastoreConfigurationGetDatastoreDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DatastoreConfigurationGetDatastoreDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this datastore configuration get datastore default response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this datastore configuration get datastore default response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this datastore configuration get datastore default response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this datastore configuration get datastore default response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this datastore configuration get datastore default response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the datastore configuration get datastore default response
func (o *DatastoreConfigurationGetDatastoreDefault) Code() int {
	return o._statusCode
}

func (o *DatastoreConfigurationGetDatastoreDefault) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] DatastoreConfiguration_GetDatastore default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreDefault) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}][%d] DatastoreConfiguration_GetDatastore default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
