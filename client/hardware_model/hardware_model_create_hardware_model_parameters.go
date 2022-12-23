// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

// NewHardwareModelCreateHardwareModelParams creates a new HardwareModelCreateHardwareModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelCreateHardwareModelParams() *HardwareModelCreateHardwareModelParams {
	return &HardwareModelCreateHardwareModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelCreateHardwareModelParamsWithTimeout creates a new HardwareModelCreateHardwareModelParams object
// with the ability to set a timeout on a request.
func NewHardwareModelCreateHardwareModelParamsWithTimeout(timeout time.Duration) *HardwareModelCreateHardwareModelParams {
	return &HardwareModelCreateHardwareModelParams{
		timeout: timeout,
	}
}

// NewHardwareModelCreateHardwareModelParamsWithContext creates a new HardwareModelCreateHardwareModelParams object
// with the ability to set a context for a request.
func NewHardwareModelCreateHardwareModelParamsWithContext(ctx context.Context) *HardwareModelCreateHardwareModelParams {
	return &HardwareModelCreateHardwareModelParams{
		Context: ctx,
	}
}

// NewHardwareModelCreateHardwareModelParamsWithHTTPClient creates a new HardwareModelCreateHardwareModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelCreateHardwareModelParamsWithHTTPClient(client *http.Client) *HardwareModelCreateHardwareModelParams {
	return &HardwareModelCreateHardwareModelParams{
		HTTPClient: client,
	}
}

/*
HardwareModelCreateHardwareModelParams contains all the parameters to send to the API endpoint

	for the hardware model create hardware model operation.

	Typically these are written to a http.Request.
*/
type HardwareModelCreateHardwareModelParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.SysModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model create hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelCreateHardwareModelParams) WithDefaults() *HardwareModelCreateHardwareModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model create hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelCreateHardwareModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) WithTimeout(timeout time.Duration) *HardwareModelCreateHardwareModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) WithContext(ctx context.Context) *HardwareModelCreateHardwareModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) WithHTTPClient(client *http.Client) *HardwareModelCreateHardwareModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) WithXRequestID(xRequestID *string) *HardwareModelCreateHardwareModelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) WithBody(body *models.SysModel) *HardwareModelCreateHardwareModelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the hardware model create hardware model params
func (o *HardwareModelCreateHardwareModelParams) SetBody(body *models.SysModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelCreateHardwareModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
