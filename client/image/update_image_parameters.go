package image

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

// UpdateParams creates a new ImageConfigurationUpdateImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateParams() *ImageConfigurationUpdateImageParams {
	return &ImageConfigurationUpdateImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationUpdateImageParamsWithTimeout creates a new ImageConfigurationUpdateImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationUpdateImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationUpdateImageParams {
	return &ImageConfigurationUpdateImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationUpdateImageParamsWithContext creates a new ImageConfigurationUpdateImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationUpdateImageParamsWithContext(ctx context.Context) *ImageConfigurationUpdateImageParams {
	return &ImageConfigurationUpdateImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationUpdateImageParamsWithHTTPClient creates a new ImageConfigurationUpdateImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationUpdateImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationUpdateImageParams {
	return &ImageConfigurationUpdateImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationUpdateImageParams contains all the parameters to send to the API endpoint

	for the image configuration update image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationUpdateImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Image

	/* ID.

	   System defined universally unique Id of the image.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration update image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUpdateImageParams) WithDefaults() *ImageConfigurationUpdateImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration update image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUpdateImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationUpdateImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithContext(ctx context.Context) *ImageConfigurationUpdateImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationUpdateImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationUpdateImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithBody(body *models.Image) *ImageConfigurationUpdateImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetBody(body *models.Image) {
	o.Body = body
}

// WithID adds the id to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) WithID(id string) *ImageConfigurationUpdateImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image configuration update image params
func (o *ImageConfigurationUpdateImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationUpdateImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
