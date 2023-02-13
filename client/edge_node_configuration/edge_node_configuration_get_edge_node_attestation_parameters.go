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

// NewEdgeNodeConfigurationGetEdgeNodeAttestationParams creates a new EdgeNodeConfigurationGetEdgeNodeAttestationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationGetEdgeNodeAttestationParams() *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithTimeout creates a new EdgeNodeConfigurationGetEdgeNodeAttestationParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithContext creates a new EdgeNodeConfigurationGetEdgeNodeAttestationParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithHTTPClient creates a new EdgeNodeConfigurationGetEdgeNodeAttestationParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationGetEdgeNodeAttestationParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationParams contains all the parameters to send to the API endpoint

	for the edge node configuration get edge node attestation operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration get edge node attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithDefaults() *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration get edge node attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WithID(id string) *EdgeNodeConfigurationGetEdgeNodeAttestationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration get edge node attestation params
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
