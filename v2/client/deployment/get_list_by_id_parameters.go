// Code generated by go-swagger; DO NOT EDIT.

package deployment

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
)

// NewGetListbyIdParams creates a new GetListbyIdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetListbyIdParams() *GetListbyIdParams {
	return &GetListbyIdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetListbyIdParamsWithTimeout creates a new GetListbyIdParams object
// with the ability to set a timeout on a request.
func NewGetListbyIdParamsWithTimeout(timeout time.Duration) *GetListbyIdParams {
	return &GetListbyIdParams{
		timeout: timeout,
	}
}

// NewGetListbyIdParamsWithContext creates a new GetListbyIdParams object
// with the ability to set a context for a request.
func NewGetListbyIdParamsWithContext(ctx context.Context) *GetListbyIdParams {
	return &GetListbyIdParams{
		Context: ctx,
	}
}

// NewGetListbyIdParamsWithHTTPClient creates a new GetListbyIdParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetListbyIdParamsWithHTTPClient(client *http.Client) *GetListbyIdParams {
	return &GetListbyIdParams{
		HTTPClient: client,
	}
}

/*
GetListbyIdParams contains all the parameters to send to the API endpoint

	for the resource group get deployment listby idv2 operation.

	Typically these are written to a http.Request.
*/
type GetListbyIdParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource group get deployment listby idv2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListbyIdParams) WithDefaults() *GetListbyIdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource group get deployment listby idv2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListbyIdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) WithTimeout(timeout time.Duration) *GetListbyIdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) WithContext(ctx context.Context) *GetListbyIdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) WithHTTPClient(client *http.Client) *GetListbyIdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) WithXRequestID(xRequestID *string) *GetListbyIdParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectID adds the projectID to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) WithProjectID(projectID string) *GetListbyIdParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the resource group get deployment listby idv2 params
func (o *GetListbyIdParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetListbyIdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
