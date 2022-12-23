// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

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

// NewEdgeNodeStatusGetEdgeNodeRawStatusParams creates a new EdgeNodeStatusGetEdgeNodeRawStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeStatusGetEdgeNodeRawStatusParams() *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithTimeout creates a new EdgeNodeStatusGetEdgeNodeRawStatusParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusParams{
		timeout: timeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithContext creates a new EdgeNodeStatusGetEdgeNodeRawStatusParams object
// with the ability to set a context for a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusParams{
		Context: ctx,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithHTTPClient creates a new EdgeNodeStatusGetEdgeNodeRawStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusParamsWithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusParams contains all the parameters to send to the API endpoint

	for the edge node status get edge node raw status operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node status get edge node raw status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithDefaults() *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node status get edge node raw status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithXRequestID(xRequestID *string) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WithID(id string) *EdgeNodeStatusGetEdgeNodeRawStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node status get edge node raw status params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeStatusGetEdgeNodeRawStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
