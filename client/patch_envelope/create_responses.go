// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// PatchEnvelopeConfigurationCreatePatchEnvelopeReader is a Reader for the PatchEnvelopeConfigurationCreatePatchEnvelope structure.
type PatchEnvelopeConfigurationCreatePatchEnvelopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationCreatePatchEnvelopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeOK creates a PatchEnvelopeConfigurationCreatePatchEnvelopeOK with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeOK() *PatchEnvelopeConfigurationCreatePatchEnvelopeOK {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeOK{}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration create patch envelope o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration create patch envelope o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration create patch envelope o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration create patch envelope o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration create patch envelope o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration create patch envelope o k response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized creates a PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized() *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized{}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration create patch envelope unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration create patch envelope unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration create patch envelope unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration create patch envelope unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration create patch envelope unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration create patch envelope unauthorized response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeForbidden creates a PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeForbidden() *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden{}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration create patch envelope forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration create patch envelope forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration create patch envelope forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration create patch envelope forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration create patch envelope forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration create patch envelope forbidden response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError creates a PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError() *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError{}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration create patch envelope internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration create patch envelope internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration create patch envelope internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration create patch envelope internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration create patch envelope internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration create patch envelope internal server error response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout creates a PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout() *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration create patch envelope gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration create patch envelope gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration create patch envelope gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration create patch envelope gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration create patch envelope gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration create patch envelope gateway timeout response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] patchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationCreatePatchEnvelopeDefault creates a PatchEnvelopeConfigurationCreatePatchEnvelopeDefault with default headers values
func NewPatchEnvelopeConfigurationCreatePatchEnvelopeDefault(code int) *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault {
	return &PatchEnvelopeConfigurationCreatePatchEnvelopeDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationCreatePatchEnvelopeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationCreatePatchEnvelopeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration create patch envelope default response has a 2xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration create patch envelope default response has a 3xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration create patch envelope default response has a 4xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration create patch envelope default response has a 5xx status code
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration create patch envelope default response a status code equal to that given
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration create patch envelope default response
func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] PatchEnvelopeConfiguration_CreatePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) String() string {
	return fmt.Sprintf("[POST /v1/patch-envelope][%d] PatchEnvelopeConfiguration_CreatePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationCreatePatchEnvelopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
