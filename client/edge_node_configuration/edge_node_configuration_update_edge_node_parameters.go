// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

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

// NewEdgeNodeConfigurationUpdateEdgeNodeParams creates a new EdgeNodeConfigurationUpdateEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationUpdateEdgeNodeParams() *EdgeNodeConfigurationUpdateEdgeNodeParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationUpdateEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationUpdateEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationUpdateEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeParams contains all the parameters to send to the API endpoint

	for the edge node configuration update edge node operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body EdgeNodeConfigurationUpdateEdgeNodeBody

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration update edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration update edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithBody(body EdgeNodeConfigurationUpdateEdgeNodeBody) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetBody(body EdgeNodeConfigurationUpdateEdgeNodeBody) {
	o.Body = body
}

// WithID adds the id to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WithID(id string) *EdgeNodeConfigurationUpdateEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration update edge node params
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationUpdateEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
