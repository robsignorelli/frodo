// Code generated by Frodo from example/names/name_service.go - DO NOT EDIT
//
//   https://github.com/monadicstack/frodo
//
package names

import (
	"context"
	"fmt"

	"github.com/monadicstack/frodo/example/names"
	"github.com/monadicstack/frodo/rpc"
)

// NewNameServiceClient creates an RPC client that conforms to the NameService interface, but delegates
// work to remote instances. You must supply the base address of the remote service gateway instance or
// the load balancer for that service.
//
// NameService performs parsing/processing on a person's name. This is primarily just
// used as a reference service for integration testing our generated clients.
func NewNameServiceClient(address string, options ...rpc.ClientOption) *NameServiceClient {
	rpcClient := rpc.NewClient("NameService", address, options...)
	rpcClient.PathPrefix = ""
	return &NameServiceClient{Client: rpcClient}
}

// NameServiceClient manages all interaction w/ a remote NameService instance by letting you invoke functions
// on this instance as if you were doing it locally (hence... RPC client). You shouldn't instantiate this
// manually. Instead, you should utilize the NewNameServiceClient() function to properly set this up.
type NameServiceClient struct {
	rpc.Client
}

// Download returns a raw CSV file containing the parsed name.
func (client *NameServiceClient) Download(ctx context.Context, request *names.DownloadRequest) (*names.DownloadResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.DownloadResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.Download", request, response)
	return response, err
}

// DownloadExt returns a raw CSV file containing the parsed name. This differs from Download
// by giving you the "Ext" knob which will let you exercise the content type and disposition
// interfaces that Frodo supports for raw responses.
func (client *NameServiceClient) DownloadExt(ctx context.Context, request *names.DownloadExtRequest) (*names.DownloadExtResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.DownloadExtResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.DownloadExt", request, response)
	return response, err
}

// FirstName extracts just the first name from a full name string.
func (client *NameServiceClient) FirstName(ctx context.Context, request *names.FirstNameRequest) (*names.FirstNameResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.FirstNameResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.FirstName", request, response)
	return response, err
}

// LastName extracts just the last name from a full name string.
func (client *NameServiceClient) LastName(ctx context.Context, request *names.LastNameRequest) (*names.LastNameResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.LastNameResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.LastName", request, response)
	return response, err
}

// SortName establishes the "phone book" name for the given full name.
func (client *NameServiceClient) SortName(ctx context.Context, request *names.SortNameRequest) (*names.SortNameResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.SortNameResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.SortName", request, response)
	return response, err
}

// Split separates a first and last name.
func (client *NameServiceClient) Split(ctx context.Context, request *names.SplitRequest) (*names.SplitResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &names.SplitResponse{}
	err := client.Invoke(ctx, "POST", "/NameService.Split", request, response)
	return response, err
}

// NameServiceProxy fully implements the NameService interface, but delegates all operations to a "real"
// instance of the service. You can embed this type in a struct of your choice so you can "override" or
// decorate operations as you see fit. Any operations on NameService that you don't explicitly define will
// simply delegate to the default implementation of the underlying 'Service' value.
//
// Since you have access to the underlying service, you are able to both implement custom handling logic AND
// call the "real" implementation, so this can be used as special middleware that applies to only certain operations.
type NameServiceProxy struct {
	Service names.NameService
}

func (proxy *NameServiceProxy) Download(ctx context.Context, request *names.DownloadRequest) (*names.DownloadResponse, error) {
	return proxy.Service.Download(ctx, request)
}

func (proxy *NameServiceProxy) DownloadExt(ctx context.Context, request *names.DownloadExtRequest) (*names.DownloadExtResponse, error) {
	return proxy.Service.DownloadExt(ctx, request)
}

func (proxy *NameServiceProxy) FirstName(ctx context.Context, request *names.FirstNameRequest) (*names.FirstNameResponse, error) {
	return proxy.Service.FirstName(ctx, request)
}

func (proxy *NameServiceProxy) LastName(ctx context.Context, request *names.LastNameRequest) (*names.LastNameResponse, error) {
	return proxy.Service.LastName(ctx, request)
}

func (proxy *NameServiceProxy) SortName(ctx context.Context, request *names.SortNameRequest) (*names.SortNameResponse, error) {
	return proxy.Service.SortName(ctx, request)
}

func (proxy *NameServiceProxy) Split(ctx context.Context, request *names.SplitRequest) (*names.SplitResponse, error) {
	return proxy.Service.Split(ctx, request)
}