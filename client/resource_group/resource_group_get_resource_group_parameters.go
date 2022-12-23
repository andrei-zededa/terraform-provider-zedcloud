// Code generated by go-swagger; DO NOT EDIT.

package resource_group

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

// NewResourceGroupGetResourceGroupParams creates a new ResourceGroupGetResourceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceGroupGetResourceGroupParams() *ResourceGroupGetResourceGroupParams {
	return &ResourceGroupGetResourceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupGetResourceGroupParamsWithTimeout creates a new ResourceGroupGetResourceGroupParams object
// with the ability to set a timeout on a request.
func NewResourceGroupGetResourceGroupParamsWithTimeout(timeout time.Duration) *ResourceGroupGetResourceGroupParams {
	return &ResourceGroupGetResourceGroupParams{
		timeout: timeout,
	}
}

// NewResourceGroupGetResourceGroupParamsWithContext creates a new ResourceGroupGetResourceGroupParams object
// with the ability to set a context for a request.
func NewResourceGroupGetResourceGroupParamsWithContext(ctx context.Context) *ResourceGroupGetResourceGroupParams {
	return &ResourceGroupGetResourceGroupParams{
		Context: ctx,
	}
}

// NewResourceGroupGetResourceGroupParamsWithHTTPClient creates a new ResourceGroupGetResourceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceGroupGetResourceGroupParamsWithHTTPClient(client *http.Client) *ResourceGroupGetResourceGroupParams {
	return &ResourceGroupGetResourceGroupParams{
		HTTPClient: client,
	}
}

/*
ResourceGroupGetResourceGroupParams contains all the parameters to send to the API endpoint

	for the resource group get resource group operation.

	Typically these are written to a http.Request.
*/
type ResourceGroupGetResourceGroupParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the resource group
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource group get resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupGetResourceGroupParams) WithDefaults() *ResourceGroupGetResourceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource group get resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupGetResourceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) WithTimeout(timeout time.Duration) *ResourceGroupGetResourceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) WithContext(ctx context.Context) *ResourceGroupGetResourceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) WithHTTPClient(client *http.Client) *ResourceGroupGetResourceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) WithXRequestID(xRequestID *string) *ResourceGroupGetResourceGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) WithID(id string) *ResourceGroupGetResourceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resource group get resource group params
func (o *ResourceGroupGetResourceGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupGetResourceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
