// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// NewIdentityAccessManagementUpdateUserParams creates a new IdentityAccessManagementUpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementUpdateUserParams() *IdentityAccessManagementUpdateUserParams {
	return &IdentityAccessManagementUpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementUpdateUserParamsWithTimeout creates a new IdentityAccessManagementUpdateUserParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementUpdateUserParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateUserParams {
	return &IdentityAccessManagementUpdateUserParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementUpdateUserParamsWithContext creates a new IdentityAccessManagementUpdateUserParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementUpdateUserParamsWithContext(ctx context.Context) *IdentityAccessManagementUpdateUserParams {
	return &IdentityAccessManagementUpdateUserParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementUpdateUserParamsWithHTTPClient creates a new IdentityAccessManagementUpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementUpdateUserParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateUserParams {
	return &IdentityAccessManagementUpdateUserParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementUpdateUserParams contains all the parameters to send to the API endpoint

	for the identity access management update user operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementUpdateUserParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.DetailedUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateUserParams) WithDefaults() *IdentityAccessManagementUpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) WithContext(ctx context.Context) *IdentityAccessManagementUpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementUpdateUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) WithBody(body *models.DetailedUser) *IdentityAccessManagementUpdateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management update user params
func (o *IdentityAccessManagementUpdateUserParams) SetBody(body *models.DetailedUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementUpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
