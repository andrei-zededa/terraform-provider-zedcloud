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

// DatastoreConfigurationQueryDatastoreProjectListReader is a Reader for the DatastoreConfigurationQueryDatastoreProjectList structure.
type DatastoreConfigurationQueryDatastoreProjectListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatastoreConfigurationQueryDatastoreProjectListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatastoreConfigurationQueryDatastoreProjectListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatastoreConfigurationQueryDatastoreProjectListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatastoreConfigurationQueryDatastoreProjectListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatastoreConfigurationQueryDatastoreProjectListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatastoreConfigurationQueryDatastoreProjectListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDatastoreConfigurationQueryDatastoreProjectListGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDatastoreConfigurationQueryDatastoreProjectListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDatastoreConfigurationQueryDatastoreProjectListOK creates a DatastoreConfigurationQueryDatastoreProjectListOK with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListOK() *DatastoreConfigurationQueryDatastoreProjectListOK {
	return &DatastoreConfigurationQueryDatastoreProjectListOK{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListOK describes a response with status code 200, with default header values.

A successful response.
*/
type DatastoreConfigurationQueryDatastoreProjectListOK struct {
	Payload *models.DatastoreProjectList
}

// IsSuccess returns true when this datastore configuration query datastore project list o k response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datastore configuration query datastore project list o k response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list o k response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastore project list o k response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastore project list o k response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datastore configuration query datastore project list o k response
func (o *DatastoreConfigurationQueryDatastoreProjectListOK) Code() int {
	return 200
}

func (o *DatastoreConfigurationQueryDatastoreProjectListOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListOK) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListOK) GetPayload() *models.DatastoreProjectList {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreProjectList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListBadRequest creates a DatastoreConfigurationQueryDatastoreProjectListBadRequest with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListBadRequest() *DatastoreConfigurationQueryDatastoreProjectListBadRequest {
	return &DatastoreConfigurationQueryDatastoreProjectListBadRequest{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type DatastoreConfigurationQueryDatastoreProjectListBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastore project list bad request response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastore project list bad request response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list bad request response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastore project list bad request response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastore project list bad request response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the datastore configuration query datastore project list bad request response
func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) Code() int {
	return 400
}

func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListUnauthorized creates a DatastoreConfigurationQueryDatastoreProjectListUnauthorized with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListUnauthorized() *DatastoreConfigurationQueryDatastoreProjectListUnauthorized {
	return &DatastoreConfigurationQueryDatastoreProjectListUnauthorized{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DatastoreConfigurationQueryDatastoreProjectListUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastore project list unauthorized response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastore project list unauthorized response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list unauthorized response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastore project list unauthorized response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastore project list unauthorized response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datastore configuration query datastore project list unauthorized response
func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) Code() int {
	return 401
}

func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListForbidden creates a DatastoreConfigurationQueryDatastoreProjectListForbidden with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListForbidden() *DatastoreConfigurationQueryDatastoreProjectListForbidden {
	return &DatastoreConfigurationQueryDatastoreProjectListForbidden{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DatastoreConfigurationQueryDatastoreProjectListForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastore project list forbidden response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastore project list forbidden response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list forbidden response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastore project list forbidden response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastore project list forbidden response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datastore configuration query datastore project list forbidden response
func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) Code() int {
	return 403
}

func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListInternalServerError creates a DatastoreConfigurationQueryDatastoreProjectListInternalServerError with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListInternalServerError() *DatastoreConfigurationQueryDatastoreProjectListInternalServerError {
	return &DatastoreConfigurationQueryDatastoreProjectListInternalServerError{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DatastoreConfigurationQueryDatastoreProjectListInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastore project list internal server error response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastore project list internal server error response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list internal server error response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastore project list internal server error response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration query datastore project list internal server error response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datastore configuration query datastore project list internal server error response
func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) Code() int {
	return 500
}

func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListGatewayTimeout creates a DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListGatewayTimeout() *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout {
	return &DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout{}
}

/*
DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastore project list gateway timeout response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastore project list gateway timeout response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastore project list gateway timeout response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastore project list gateway timeout response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration query datastore project list gateway timeout response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the datastore configuration query datastore project list gateway timeout response
func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) Code() int {
	return 504
}

func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] datastoreConfigurationQueryDatastoreProjectListGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoreProjectListDefault creates a DatastoreConfigurationQueryDatastoreProjectListDefault with default headers values
func NewDatastoreConfigurationQueryDatastoreProjectListDefault(code int) *DatastoreConfigurationQueryDatastoreProjectListDefault {
	return &DatastoreConfigurationQueryDatastoreProjectListDefault{
		_statusCode: code,
	}
}

/*
DatastoreConfigurationQueryDatastoreProjectListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DatastoreConfigurationQueryDatastoreProjectListDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this datastore configuration query datastore project list default response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this datastore configuration query datastore project list default response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this datastore configuration query datastore project list default response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this datastore configuration query datastore project list default response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this datastore configuration query datastore project list default response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the datastore configuration query datastore project list default response
func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) Code() int {
	return o._statusCode
}

func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] DatastoreConfiguration_QueryDatastoreProjectList default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) String() string {
	return fmt.Sprintf("[GET /v1/datastores/id/{id}/projects][%d] DatastoreConfiguration_QueryDatastoreProjectList default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoreProjectListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
