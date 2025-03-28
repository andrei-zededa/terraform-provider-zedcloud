// Code generated by go-swagger; DO NOT EDIT.

package profile_deployment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new profile deployment service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new profile deployment service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new profile deployment service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for profile deployment service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProfileDeploymentServiceCreateProfileDeployment(params *ProfileDeploymentServiceCreateProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceCreateProfileDeploymentOK, error)

	ProfileDeploymentServiceDeleteProfileDeployment(params *ProfileDeploymentServiceDeleteProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceDeleteProfileDeploymentOK, error)

	ProfileDeploymentServiceGetProfileDeployment(params *ProfileDeploymentServiceGetProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceGetProfileDeploymentOK, error)

	ProfileDeploymentServiceGetProfileDeploymentByName(params *ProfileDeploymentServiceGetProfileDeploymentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceGetProfileDeploymentByNameOK, error)

	ProfileDeploymentServiceQueryProfileDeploymentResourceStatus(params *ProfileDeploymentServiceQueryProfileDeploymentResourceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceQueryProfileDeploymentResourceStatusOK, error)

	ProfileDeploymentServiceQueryProfileDeployments(params *ProfileDeploymentServiceQueryProfileDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceQueryProfileDeploymentsOK, error)

	ProfileDeploymentServiceUpdateProfileDeployment(params *ProfileDeploymentServiceUpdateProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceUpdateProfileDeploymentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ProfileDeploymentServiceCreateProfileDeployment creates profile deployment

Create a profile deployment record.
*/
func (a *Client) ProfileDeploymentServiceCreateProfileDeployment(params *ProfileDeploymentServiceCreateProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceCreateProfileDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceCreateProfileDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_CreateProfileDeployment",
		Method:             "POST",
		PathPattern:        "/v1/profiledeployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceCreateProfileDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceCreateProfileDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceCreateProfileDeploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceDeleteProfileDeployment deletes profile deployment

Delete a profile deployment record.
*/
func (a *Client) ProfileDeploymentServiceDeleteProfileDeployment(params *ProfileDeploymentServiceDeleteProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceDeleteProfileDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceDeleteProfileDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_DeleteProfileDeployment",
		Method:             "DELETE",
		PathPattern:        "/v1/profiledeployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceDeleteProfileDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceDeleteProfileDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceDeleteProfileDeploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceGetProfileDeployment gets profile deployment

Get the configuration (without security details) of a profile deployment record.
*/
func (a *Client) ProfileDeploymentServiceGetProfileDeployment(params *ProfileDeploymentServiceGetProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceGetProfileDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceGetProfileDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_GetProfileDeployment",
		Method:             "GET",
		PathPattern:        "/v1/profiledeployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceGetProfileDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceGetProfileDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceGetProfileDeploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceGetProfileDeploymentByName gets profile deployment by name

Get the configuration (without security details) of a profile deployment record by name.
*/
func (a *Client) ProfileDeploymentServiceGetProfileDeploymentByName(params *ProfileDeploymentServiceGetProfileDeploymentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceGetProfileDeploymentByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceGetProfileDeploymentByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_GetProfileDeploymentByName",
		Method:             "GET",
		PathPattern:        "/v1/profiledeployments/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceGetProfileDeploymentByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceGetProfileDeploymentByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceGetProfileDeploymentByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceQueryProfileDeploymentResourceStatus queries profile deployments resource status

Query the profile deployment resourcs status.
*/
func (a *Client) ProfileDeploymentServiceQueryProfileDeploymentResourceStatus(params *ProfileDeploymentServiceQueryProfileDeploymentResourceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceQueryProfileDeploymentResourceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceQueryProfileDeploymentResourceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_QueryProfileDeploymentResourceStatus",
		Method:             "GET",
		PathPattern:        "/v1/profiledeployments/id/{id}/resourcestatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceQueryProfileDeploymentResourceStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceQueryProfileDeploymentResourceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceQueryProfileDeploymentResourceStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceQueryProfileDeployments queries profile deployments

Query the profile deployment records.
*/
func (a *Client) ProfileDeploymentServiceQueryProfileDeployments(params *ProfileDeploymentServiceQueryProfileDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceQueryProfileDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceQueryProfileDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_QueryProfileDeployments",
		Method:             "GET",
		PathPattern:        "/v1/profiledeployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceQueryProfileDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceQueryProfileDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceQueryProfileDeploymentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProfileDeploymentServiceUpdateProfileDeployment updates profile deployment

Update a profile deployment record. The usual pattern to update a profile deployment record is to retrieve the record and update with the modified values in a new body to update the profile deployment record.
*/
func (a *Client) ProfileDeploymentServiceUpdateProfileDeployment(params *ProfileDeploymentServiceUpdateProfileDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProfileDeploymentServiceUpdateProfileDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileDeploymentServiceUpdateProfileDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProfileDeploymentService_UpdateProfileDeployment",
		Method:             "PUT",
		PathPattern:        "/v1/profiledeployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileDeploymentServiceUpdateProfileDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProfileDeploymentServiceUpdateProfileDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProfileDeploymentServiceUpdateProfileDeploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
