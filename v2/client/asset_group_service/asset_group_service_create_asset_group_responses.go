// Code generated by go-swagger; DO NOT EDIT.

package asset_group_service

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

// AssetGroupServiceCreateAssetGroupReader is a Reader for the AssetGroupServiceCreateAssetGroup structure.
type AssetGroupServiceCreateAssetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetGroupServiceCreateAssetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetGroupServiceCreateAssetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssetGroupServiceCreateAssetGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssetGroupServiceCreateAssetGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssetGroupServiceCreateAssetGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAssetGroupServiceCreateAssetGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssetGroupServiceCreateAssetGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewAssetGroupServiceCreateAssetGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAssetGroupServiceCreateAssetGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetGroupServiceCreateAssetGroupOK creates a AssetGroupServiceCreateAssetGroupOK with default headers values
func NewAssetGroupServiceCreateAssetGroupOK() *AssetGroupServiceCreateAssetGroupOK {
	return &AssetGroupServiceCreateAssetGroupOK{}
}

/*
AssetGroupServiceCreateAssetGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetGroupServiceCreateAssetGroupOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group o k response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this asset group service create asset group o k response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group o k response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service create asset group o k response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service create asset group o k response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the asset group service create asset group o k response
func (o *AssetGroupServiceCreateAssetGroupOK) Code() int {
	return 200
}

func (o *AssetGroupServiceCreateAssetGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupOK %s", 200, payload)
}

func (o *AssetGroupServiceCreateAssetGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupOK %s", 200, payload)
}

func (o *AssetGroupServiceCreateAssetGroupOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupBadRequest creates a AssetGroupServiceCreateAssetGroupBadRequest with default headers values
func NewAssetGroupServiceCreateAssetGroupBadRequest() *AssetGroupServiceCreateAssetGroupBadRequest {
	return &AssetGroupServiceCreateAssetGroupBadRequest{}
}

/*
AssetGroupServiceCreateAssetGroupBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type AssetGroupServiceCreateAssetGroupBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group bad request response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group bad request response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group bad request response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service create asset group bad request response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service create asset group bad request response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the asset group service create asset group bad request response
func (o *AssetGroupServiceCreateAssetGroupBadRequest) Code() int {
	return 400
}

func (o *AssetGroupServiceCreateAssetGroupBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupBadRequest %s", 400, payload)
}

func (o *AssetGroupServiceCreateAssetGroupBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupBadRequest %s", 400, payload)
}

func (o *AssetGroupServiceCreateAssetGroupBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupUnauthorized creates a AssetGroupServiceCreateAssetGroupUnauthorized with default headers values
func NewAssetGroupServiceCreateAssetGroupUnauthorized() *AssetGroupServiceCreateAssetGroupUnauthorized {
	return &AssetGroupServiceCreateAssetGroupUnauthorized{}
}

/*
AssetGroupServiceCreateAssetGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type AssetGroupServiceCreateAssetGroupUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group unauthorized response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group unauthorized response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group unauthorized response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service create asset group unauthorized response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service create asset group unauthorized response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the asset group service create asset group unauthorized response
func (o *AssetGroupServiceCreateAssetGroupUnauthorized) Code() int {
	return 401
}

func (o *AssetGroupServiceCreateAssetGroupUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupUnauthorized %s", 401, payload)
}

func (o *AssetGroupServiceCreateAssetGroupUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupUnauthorized %s", 401, payload)
}

func (o *AssetGroupServiceCreateAssetGroupUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupForbidden creates a AssetGroupServiceCreateAssetGroupForbidden with default headers values
func NewAssetGroupServiceCreateAssetGroupForbidden() *AssetGroupServiceCreateAssetGroupForbidden {
	return &AssetGroupServiceCreateAssetGroupForbidden{}
}

/*
AssetGroupServiceCreateAssetGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type AssetGroupServiceCreateAssetGroupForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group forbidden response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group forbidden response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group forbidden response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service create asset group forbidden response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service create asset group forbidden response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the asset group service create asset group forbidden response
func (o *AssetGroupServiceCreateAssetGroupForbidden) Code() int {
	return 403
}

func (o *AssetGroupServiceCreateAssetGroupForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupForbidden %s", 403, payload)
}

func (o *AssetGroupServiceCreateAssetGroupForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupForbidden %s", 403, payload)
}

func (o *AssetGroupServiceCreateAssetGroupForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupConflict creates a AssetGroupServiceCreateAssetGroupConflict with default headers values
func NewAssetGroupServiceCreateAssetGroupConflict() *AssetGroupServiceCreateAssetGroupConflict {
	return &AssetGroupServiceCreateAssetGroupConflict{}
}

/*
AssetGroupServiceCreateAssetGroupConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this asset group record will conflict with an already existing asset group record.
*/
type AssetGroupServiceCreateAssetGroupConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group conflict response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group conflict response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group conflict response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service create asset group conflict response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service create asset group conflict response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the asset group service create asset group conflict response
func (o *AssetGroupServiceCreateAssetGroupConflict) Code() int {
	return 409
}

func (o *AssetGroupServiceCreateAssetGroupConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupConflict %s", 409, payload)
}

func (o *AssetGroupServiceCreateAssetGroupConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupConflict %s", 409, payload)
}

func (o *AssetGroupServiceCreateAssetGroupConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupInternalServerError creates a AssetGroupServiceCreateAssetGroupInternalServerError with default headers values
func NewAssetGroupServiceCreateAssetGroupInternalServerError() *AssetGroupServiceCreateAssetGroupInternalServerError {
	return &AssetGroupServiceCreateAssetGroupInternalServerError{}
}

/*
AssetGroupServiceCreateAssetGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type AssetGroupServiceCreateAssetGroupInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group internal server error response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group internal server error response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group internal server error response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service create asset group internal server error response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this asset group service create asset group internal server error response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the asset group service create asset group internal server error response
func (o *AssetGroupServiceCreateAssetGroupInternalServerError) Code() int {
	return 500
}

func (o *AssetGroupServiceCreateAssetGroupInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupInternalServerError %s", 500, payload)
}

func (o *AssetGroupServiceCreateAssetGroupInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupInternalServerError %s", 500, payload)
}

func (o *AssetGroupServiceCreateAssetGroupInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupGatewayTimeout creates a AssetGroupServiceCreateAssetGroupGatewayTimeout with default headers values
func NewAssetGroupServiceCreateAssetGroupGatewayTimeout() *AssetGroupServiceCreateAssetGroupGatewayTimeout {
	return &AssetGroupServiceCreateAssetGroupGatewayTimeout{}
}

/*
AssetGroupServiceCreateAssetGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type AssetGroupServiceCreateAssetGroupGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service create asset group gateway timeout response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service create asset group gateway timeout response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service create asset group gateway timeout response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service create asset group gateway timeout response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this asset group service create asset group gateway timeout response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the asset group service create asset group gateway timeout response
func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) Code() int {
	return 504
}

func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupGatewayTimeout %s", 504, payload)
}

func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] assetGroupServiceCreateAssetGroupGatewayTimeout %s", 504, payload)
}

func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceCreateAssetGroupDefault creates a AssetGroupServiceCreateAssetGroupDefault with default headers values
func NewAssetGroupServiceCreateAssetGroupDefault(code int) *AssetGroupServiceCreateAssetGroupDefault {
	return &AssetGroupServiceCreateAssetGroupDefault{
		_statusCode: code,
	}
}

/*
AssetGroupServiceCreateAssetGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetGroupServiceCreateAssetGroupDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this asset group service create asset group default response has a 2xx status code
func (o *AssetGroupServiceCreateAssetGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this asset group service create asset group default response has a 3xx status code
func (o *AssetGroupServiceCreateAssetGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this asset group service create asset group default response has a 4xx status code
func (o *AssetGroupServiceCreateAssetGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this asset group service create asset group default response has a 5xx status code
func (o *AssetGroupServiceCreateAssetGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this asset group service create asset group default response a status code equal to that given
func (o *AssetGroupServiceCreateAssetGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the asset group service create asset group default response
func (o *AssetGroupServiceCreateAssetGroupDefault) Code() int {
	return o._statusCode
}

func (o *AssetGroupServiceCreateAssetGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] AssetGroupService_CreateAssetGroup default %s", o._statusCode, payload)
}

func (o *AssetGroupServiceCreateAssetGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/assetgroups][%d] AssetGroupService_CreateAssetGroup default %s", o._statusCode, payload)
}

func (o *AssetGroupServiceCreateAssetGroupDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AssetGroupServiceCreateAssetGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
