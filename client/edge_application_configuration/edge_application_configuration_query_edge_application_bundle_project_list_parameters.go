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
	"github.com/go-openapi/swag"
)

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams creates a new EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams() *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithTimeout creates a new EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithContext creates a new EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithHTTPClient creates a new EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams contains all the parameters to send to the API endpoint

	for the edge application configuration query edge application bundle project list operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the app bundle
	*/
	ID string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy []string

	/* NextPageNum.

	   Page Number

	   Format: int64
	*/
	NextPageNum *int64

	/* NextPageSize.

	   Defines the page size

	   Format: int64
	*/
	NextPageSize *int64

	/* NextPageToken.

	   Page Token
	*/
	NextPageToken *string

	/* NextTotalPages.

	   Total number of pages to be fetched.

	   Format: int64
	*/
	NextTotalPages *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application configuration query edge application bundle project list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithDefaults() *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration query edge application bundle project list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithID(id string) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetID(id string) {
	o.ID = id
}

// WithNextOrderBy adds the nextOrderBy to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithNextOrderBy(nextOrderBy []string) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithNextPageNum(nextPageNum *int64) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithNextPageSize(nextPageSize *int64) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithNextPageToken(nextPageToken *string) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WithNextTotalPages(nextTotalPages *int64) *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the edge application configuration query edge application bundle project list params
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NextOrderBy != nil {

		// binding items for next.orderBy
		joinedNextOrderBy := o.bindParamNextOrderBy(reg)

		// query array param next.orderBy
		if err := r.SetQueryParam("next.orderBy", joinedNextOrderBy...); err != nil {
			return err
		}
	}

	if o.NextPageNum != nil {

		// query param next.pageNum
		var qrNextPageNum int64

		if o.NextPageNum != nil {
			qrNextPageNum = *o.NextPageNum
		}
		qNextPageNum := swag.FormatInt64(qrNextPageNum)
		if qNextPageNum != "" {

			if err := r.SetQueryParam("next.pageNum", qNextPageNum); err != nil {
				return err
			}
		}
	}

	if o.NextPageSize != nil {

		// query param next.pageSize
		var qrNextPageSize int64

		if o.NextPageSize != nil {
			qrNextPageSize = *o.NextPageSize
		}
		qNextPageSize := swag.FormatInt64(qrNextPageSize)
		if qNextPageSize != "" {

			if err := r.SetQueryParam("next.pageSize", qNextPageSize); err != nil {
				return err
			}
		}
	}

	if o.NextPageToken != nil {

		// query param next.pageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("next.pageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.NextTotalPages != nil {

		// query param next.totalPages
		var qrNextTotalPages int64

		if o.NextTotalPages != nil {
			qrNextTotalPages = *o.NextTotalPages
		}
		qNextTotalPages := swag.FormatInt64(qrNextTotalPages)
		if qNextTotalPages != "" {

			if err := r.SetQueryParam("next.totalPages", qNextTotalPages); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEdgeApplicationConfigurationQueryEdgeApplicationBundleProjectList binds the parameter next.orderBy
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundleProjectListParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
	nextOrderByIR := o.NextOrderBy

	var nextOrderByIC []string
	for _, nextOrderByIIR := range nextOrderByIR { // explode []string

		nextOrderByIIV := nextOrderByIIR // string as string
		nextOrderByIC = append(nextOrderByIC, nextOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	nextOrderByIS := swag.JoinByFormat(nextOrderByIC, "multi")

	return nextOrderByIS
}
