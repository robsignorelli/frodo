// Code generated by Frodo from calc/calculator_service.go - DO NOT EDIT
//
//   https://github.com/monadicstack/frodo
//
package calc

import (
	"context"
	"fmt"

	"github.com/monadicstack/frodo/example/basic/calc"
	"github.com/monadicstack/frodo/rpc"
)

// NewCalculatorServiceClient creates an RPC client that conforms to the CalculatorService interface, but delegates
// work to remote instances. You must supply the base address of the remote service gateway instance or
// the load balancer for that service.
//
// CalculatorService provides the ability to add and subtract at WEB SCALE!
func NewCalculatorServiceClient(address string, options ...rpc.ClientOption) *CalculatorServiceClient {
	rpcClient := rpc.NewClient("CalculatorService", address, options...)
	rpcClient.PathPrefix = ""
	return &CalculatorServiceClient{Client: rpcClient}
}

// CalculatorServiceClient manages all interaction w/ a remote CalculatorService instance by letting you invoke functions
// on this instance as if you were doing it locally (hence... RPC client). You shouldn't instantiate this
// manually. Instead, you should utilize the NewCalculatorServiceClient() function to properly set this up.
type CalculatorServiceClient struct {
	rpc.Client
}

// Add accepts two integers and returns a result w/ their sum.
func (client *CalculatorServiceClient) Add(ctx context.Context, request *calc.AddRequest) (*calc.AddResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &calc.AddResponse{}
	err := client.Invoke(ctx, "POST", "/CalculatorService.Add", request, response)
	return response, err
}

// Sub accepts two integers and returns a result w/ their difference.
func (client *CalculatorServiceClient) Sub(ctx context.Context, request *calc.SubRequest) (*calc.SubResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &calc.SubResponse{}
	err := client.Invoke(ctx, "GET", "/sub/:A/:B/please", request, response)
	return response, err
}

// CalculatorServiceProxy fully implements the CalculatorService interface, but delegates all operations to a "real"
// instance of the service. You can embed this type in a struct of your choice so you can "override" or
// decorate operations as you see fit. Any operations on CalculatorService that you don't explicitly define will
// simply delegate to the default implementation of the underlying 'Service' value.
//
// Since you have access to the underlying service, you are able to both implement custom handling logic AND
// call the "real" implementation, so this can be used as special middleware that applies to only certain operations.
type CalculatorServiceProxy struct {
	Service calc.CalculatorService
}

func (proxy *CalculatorServiceProxy) Add(ctx context.Context, request *calc.AddRequest) (*calc.AddResponse, error) {
	return proxy.Service.Add(ctx, request)
}

func (proxy *CalculatorServiceProxy) Sub(ctx context.Context, request *calc.SubRequest) (*calc.SubResponse, error) {
	return proxy.Service.Sub(ctx, request)
}
