// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package appconnections

import (
	"context"
	"fmt"
	"math"
	"net/url"

	appconnectionspb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListAppConnections    []gax.CallOption
	GetAppConnection      []gax.CallOption
	CreateAppConnection   []gax.CallOption
	UpdateAppConnection   []gax.CallOption
	DeleteAppConnection   []gax.CallOption
	ResolveAppConnections []gax.CallOption
	GetLocation           []gax.CallOption
	ListLocations         []gax.CallOption
	GetIamPolicy          []gax.CallOption
	SetIamPolicy          []gax.CallOption
	TestIamPermissions    []gax.CallOption
	CancelOperation       []gax.CallOption
	DeleteOperation       []gax.CallOption
	GetOperation          []gax.CallOption
	ListOperations        []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("beyondcorp.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("beyondcorp.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://beyondcorp.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListAppConnections:    []gax.CallOption{},
		GetAppConnection:      []gax.CallOption{},
		CreateAppConnection:   []gax.CallOption{},
		UpdateAppConnection:   []gax.CallOption{},
		DeleteAppConnection:   []gax.CallOption{},
		ResolveAppConnections: []gax.CallOption{},
		GetLocation:           []gax.CallOption{},
		ListLocations:         []gax.CallOption{},
		GetIamPolicy:          []gax.CallOption{},
		SetIamPolicy:          []gax.CallOption{},
		TestIamPermissions:    []gax.CallOption{},
		CancelOperation:       []gax.CallOption{},
		DeleteOperation:       []gax.CallOption{},
		GetOperation:          []gax.CallOption{},
		ListOperations:        []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from BeyondCorp API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAppConnections(context.Context, *appconnectionspb.ListAppConnectionsRequest, ...gax.CallOption) *AppConnectionIterator
	GetAppConnection(context.Context, *appconnectionspb.GetAppConnectionRequest, ...gax.CallOption) (*appconnectionspb.AppConnection, error)
	CreateAppConnection(context.Context, *appconnectionspb.CreateAppConnectionRequest, ...gax.CallOption) (*CreateAppConnectionOperation, error)
	CreateAppConnectionOperation(name string) *CreateAppConnectionOperation
	UpdateAppConnection(context.Context, *appconnectionspb.UpdateAppConnectionRequest, ...gax.CallOption) (*UpdateAppConnectionOperation, error)
	UpdateAppConnectionOperation(name string) *UpdateAppConnectionOperation
	DeleteAppConnection(context.Context, *appconnectionspb.DeleteAppConnectionRequest, ...gax.CallOption) (*DeleteAppConnectionOperation, error)
	DeleteAppConnectionOperation(name string) *DeleteAppConnectionOperation
	ResolveAppConnections(context.Context, *appconnectionspb.ResolveAppConnectionsRequest, ...gax.CallOption) *ResolveAppConnectionsResponse_AppConnectionDetailsIterator
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// Client is a client for interacting with BeyondCorp API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API Overview:
//
// The beyondcorp.googleapis.com service implements the Google Cloud
// BeyondCorp API.
//
// Data Model:
//
// The AppConnectionsService exposes the following resources:
//
//	AppConnections, named as follows:
//	projects/{project_id}/locations/{location_id}/appConnections/{app_connection_id}.
//
// The AppConnectionsService service provides methods to manage
// (create/read/update/delete) BeyondCorp AppConnections.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAppConnections lists AppConnections in a given project and location.
func (c *Client) ListAppConnections(ctx context.Context, req *appconnectionspb.ListAppConnectionsRequest, opts ...gax.CallOption) *AppConnectionIterator {
	return c.internalClient.ListAppConnections(ctx, req, opts...)
}

// GetAppConnection gets details of a single AppConnection.
func (c *Client) GetAppConnection(ctx context.Context, req *appconnectionspb.GetAppConnectionRequest, opts ...gax.CallOption) (*appconnectionspb.AppConnection, error) {
	return c.internalClient.GetAppConnection(ctx, req, opts...)
}

// CreateAppConnection creates a new AppConnection in a given project and location.
func (c *Client) CreateAppConnection(ctx context.Context, req *appconnectionspb.CreateAppConnectionRequest, opts ...gax.CallOption) (*CreateAppConnectionOperation, error) {
	return c.internalClient.CreateAppConnection(ctx, req, opts...)
}

// CreateAppConnectionOperation returns a new CreateAppConnectionOperation from a given name.
// The name must be that of a previously created CreateAppConnectionOperation, possibly from a different process.
func (c *Client) CreateAppConnectionOperation(name string) *CreateAppConnectionOperation {
	return c.internalClient.CreateAppConnectionOperation(name)
}

// UpdateAppConnection updates the parameters of a single AppConnection.
func (c *Client) UpdateAppConnection(ctx context.Context, req *appconnectionspb.UpdateAppConnectionRequest, opts ...gax.CallOption) (*UpdateAppConnectionOperation, error) {
	return c.internalClient.UpdateAppConnection(ctx, req, opts...)
}

// UpdateAppConnectionOperation returns a new UpdateAppConnectionOperation from a given name.
// The name must be that of a previously created UpdateAppConnectionOperation, possibly from a different process.
func (c *Client) UpdateAppConnectionOperation(name string) *UpdateAppConnectionOperation {
	return c.internalClient.UpdateAppConnectionOperation(name)
}

// DeleteAppConnection deletes a single AppConnection.
func (c *Client) DeleteAppConnection(ctx context.Context, req *appconnectionspb.DeleteAppConnectionRequest, opts ...gax.CallOption) (*DeleteAppConnectionOperation, error) {
	return c.internalClient.DeleteAppConnection(ctx, req, opts...)
}

// DeleteAppConnectionOperation returns a new DeleteAppConnectionOperation from a given name.
// The name must be that of a previously created DeleteAppConnectionOperation, possibly from a different process.
func (c *Client) DeleteAppConnectionOperation(name string) *DeleteAppConnectionOperation {
	return c.internalClient.DeleteAppConnectionOperation(name)
}

// ResolveAppConnections resolves AppConnections details for a given AppConnector.
// An internal method called by a connector to find AppConnections to connect
// to.
func (c *Client) ResolveAppConnections(ctx context.Context, req *appconnectionspb.ResolveAppConnectionsRequest, opts ...gax.CallOption) *ResolveAppConnectionsResponse_AppConnectionDetailsIterator {
	return c.internalClient.ResolveAppConnections(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *Client) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *Client) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *Client) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *Client) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *Client) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *Client) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// gRPCClient is a client for interacting with BeyondCorp API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client appconnectionspb.AppConnectionsServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new app connections service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API Overview:
//
// The beyondcorp.googleapis.com service implements the Google Cloud
// BeyondCorp API.
//
// Data Model:
//
// The AppConnectionsService exposes the following resources:
//
//	AppConnections, named as follows:
//	projects/{project_id}/locations/{location_id}/appConnections/{app_connection_id}.
//
// The AppConnectionsService service provides methods to manage
// (create/read/update/delete) BeyondCorp AppConnections.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		client:           appconnectionspb.NewAppConnectionsServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:  iampb.NewIAMPolicyClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListAppConnections(ctx context.Context, req *appconnectionspb.ListAppConnectionsRequest, opts ...gax.CallOption) *AppConnectionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAppConnections[0:len((*c.CallOptions).ListAppConnections):len((*c.CallOptions).ListAppConnections)], opts...)
	it := &AppConnectionIterator{}
	req = proto.Clone(req).(*appconnectionspb.ListAppConnectionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appconnectionspb.AppConnection, string, error) {
		resp := &appconnectionspb.ListAppConnectionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListAppConnections(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAppConnections(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetAppConnection(ctx context.Context, req *appconnectionspb.GetAppConnectionRequest, opts ...gax.CallOption) (*appconnectionspb.AppConnection, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAppConnection[0:len((*c.CallOptions).GetAppConnection):len((*c.CallOptions).GetAppConnection)], opts...)
	var resp *appconnectionspb.AppConnection
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAppConnection(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateAppConnection(ctx context.Context, req *appconnectionspb.CreateAppConnectionRequest, opts ...gax.CallOption) (*CreateAppConnectionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateAppConnection[0:len((*c.CallOptions).CreateAppConnection):len((*c.CallOptions).CreateAppConnection)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAppConnection(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UpdateAppConnection(ctx context.Context, req *appconnectionspb.UpdateAppConnectionRequest, opts ...gax.CallOption) (*UpdateAppConnectionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "app_connection.name", url.QueryEscape(req.GetAppConnection().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAppConnection[0:len((*c.CallOptions).UpdateAppConnection):len((*c.CallOptions).UpdateAppConnection)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateAppConnection(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteAppConnection(ctx context.Context, req *appconnectionspb.DeleteAppConnectionRequest, opts ...gax.CallOption) (*DeleteAppConnectionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteAppConnection[0:len((*c.CallOptions).DeleteAppConnection):len((*c.CallOptions).DeleteAppConnection)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteAppConnection(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ResolveAppConnections(ctx context.Context, req *appconnectionspb.ResolveAppConnectionsRequest, opts ...gax.CallOption) *ResolveAppConnectionsResponse_AppConnectionDetailsIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ResolveAppConnections[0:len((*c.CallOptions).ResolveAppConnections):len((*c.CallOptions).ResolveAppConnections)], opts...)
	it := &ResolveAppConnectionsResponse_AppConnectionDetailsIterator{}
	req = proto.Clone(req).(*appconnectionspb.ResolveAppConnectionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appconnectionspb.ResolveAppConnectionsResponse_AppConnectionDetails, string, error) {
		resp := &appconnectionspb.ResolveAppConnectionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ResolveAppConnections(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAppConnectionDetails(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateAppConnectionOperation returns a new CreateAppConnectionOperation from a given name.
// The name must be that of a previously created CreateAppConnectionOperation, possibly from a different process.
func (c *gRPCClient) CreateAppConnectionOperation(name string) *CreateAppConnectionOperation {
	return &CreateAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteAppConnectionOperation returns a new DeleteAppConnectionOperation from a given name.
// The name must be that of a previously created DeleteAppConnectionOperation, possibly from a different process.
func (c *gRPCClient) DeleteAppConnectionOperation(name string) *DeleteAppConnectionOperation {
	return &DeleteAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateAppConnectionOperation returns a new UpdateAppConnectionOperation from a given name.
// The name must be that of a previously created UpdateAppConnectionOperation, possibly from a different process.
func (c *gRPCClient) UpdateAppConnectionOperation(name string) *UpdateAppConnectionOperation {
	return &UpdateAppConnectionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
