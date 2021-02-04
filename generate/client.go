package generate

// TemplateClientGo is the text template for a service's Go client code. It creates the constructor
// function to create the client such as 'NewFooServiceClient()' and all of the types/functions to
// properly implement the service as an RPC client.
var TemplateClientGo = parseArtifactTemplate("client.go", `// !!!!!!! DO NOT EDIT !!!!!!!
// Auto-generated client code from {{ .Path }}
// !!!!!!! DO NOT EDIT !!!!!!!
package {{ .OutputPackage.Name }}

import (
	"context"
	"fmt"

	"github.com/robsignorelli/frodo/rpc"
	"{{ .Package.Import }}"
)

{{ $ctx := . }}
{{ range .Services }}
// New{{ .Name }}Client creates an RPC client that conforms to the {{ .Name }} interface, but delegates
// work to remote instances. You must supply the base address of the remote service gateway instance or
// the load balancer for that service. {{ range .Documentation }}
// {{ . }}{{end}}
func New{{ .Name }}Client(address string, options ...rpc.ClientOption) *{{ .Name }}Client {
	rpcClient := rpc.NewClient("{{ .Name }}", address, options...)
	rpcClient.PathPrefix = "{{ .HTTPPathPrefix }}"
	return &{{ .Name }}Client{Client: rpcClient} 
}

// {{ .Name }}Client manages all interaction w/ a remote {{ .Name }} instance by letting you invoke functions
// on this instance as if you were doing it locally (hence... RPC client). You shouldn't instantiate this
// manually. Instead, you should utilize the New{{ .Name }}Client() function to properly set this up.
type {{ .Name }}Client struct {
	rpc.Client
}

{{ $service := . }}
{{ range .Functions }}
{{ range .Documentation }}
// {{ . }}{{ end }}
func (client *{{ $service.Name }}Client) {{ .Name }} (ctx context.Context, request *{{ $ctx.Package.Name }}.{{ .Request.Name | NonPointer }}) (*{{ $ctx.Package.Name }}.{{ .Response.Name | NonPointer }}, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &{{ $ctx.Package.Name }}.{{ .Response.Name }}{}
	err := client.Invoke(ctx, "{{ .HTTPMethod }}", "{{ .HTTPPath }}", request, response)
	return response, err
}
{{ end }}

// {{ .Name }}Proxy fully implements the {{ .Name }} interface, but delegates all operations to a "real"
// instance of the service. You can embed this type in a struct of your choice so you can "override" or
// decorate operations as you see fit. Any operations on {{ .Name }} that you don't explicitly define will
// simply delegate to the default implementation of the underlying 'Service' value.
//
// Since you have access to the underlying service, you are able to both implement custom handling logic AND
// call the "real" implementation, so this can be used as special middleware that applies to only certain operations.
type {{ .Name }}Proxy struct {
	Service {{ $ctx.Package.Name }}.{{ .Name }}
}

{{ range .Functions }}
func (proxy *{{ $service.Name }}Proxy) {{ .Name }} (ctx context.Context, request *{{ $ctx.Package.Name }}.{{ .Request.Name }}) (*{{ $ctx.Package.Name }}.{{ .Response.Name }}, error) {
	return proxy.Service.{{ .Name }}(ctx, request)
}
{{ end }}
{{ end }}
`)
