// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

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

	"github.com/zededa/terraform-provider/models"
)

// CreateParams creates a new EdgeApplicationConfigurationCreateEdgeApplicationBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateParams() *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationCreateEdgeApplicationBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithTimeout creates a new EdgeApplicationConfigurationCreateEdgeApplicationBundleParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationCreateEdgeApplicationBundleParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithContext creates a new EdgeApplicationConfigurationCreateEdgeApplicationBundleParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationCreateEdgeApplicationBundleParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithHTTPClient creates a new EdgeApplicationConfigurationCreateEdgeApplicationBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationCreateEdgeApplicationBundleParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationCreateEdgeApplicationBundleParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationCreateEdgeApplicationBundleParams contains all the parameters to send to the API endpoint

	for the edge application configuration create edge application bundle operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationCreateEdgeApplicationBundleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Application

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application configuration create edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithDefaults() *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration create edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WithBody(body *models.Application) *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge application configuration create edge application bundle params
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) SetBody(body *models.Application) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationCreateEdgeApplicationBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
