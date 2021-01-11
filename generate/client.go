package generate

import (
	"text/template"
)

// Once Go 1.16 comes out and we can embed files in the Go binary, I should pull this out
// into a separate template file and just embed that in the binary fs.
var TemplateClientGo = template.Must(template.New("client.go").Parse(`// !!!!!!! DO NOT EDIT !!!!!!!
// Auto-generated client code from {{ .Path }}
// !!!!!!! DO NOT EDIT !!!!!!!
package {{ .Package }}

import (
	"context"
	"fmt"

	"github.com/robsignorelli/frodo/rpc"
)

{{ $ctx := . }}
{{ range .Services }}
func New{{ .Name }}Client(address string, options ...rpc.ClientOption) *{{.Name}}Client {
	fmt.Println(">>>> Creating client")
	return &{{.Name}}Client{
		Client: rpc.NewClient("{{ .Name }}", address, options...),
	}
}

type {{ .Name }}Client struct {
	rpc.Client
}

{{ $service := . }}
{{ range .Methods }}
func (client *{{ $service.Name }}Client) {{ .Name }} (ctx context.Context, request *{{ .Request.Name }}) (*{{ .Response.Name }}, error) {
	if ctx == nil {
		return nil, fmt.Errorf("precondition failed: nil context")
	}
	if request == nil {
		return nil, fmt.Errorf("precondition failed: nil request")
	}

	response := &{{ .Response.Name }}{}
	err := client.Invoke(ctx, "{{ .HTTPMethod }}", "{{ .HTTPPath }}", request, response)
	return response, err
}
{{ end }}

type {{ .Name }}Proxy struct {
	Service {{ .Name }}
}

{{ range .Methods }}
func (proxy *{{ $service.Name }}Proxy) {{ .Name }} (ctx context.Context, request *{{ .Request.Name }}) (*{{ .Response.Name }}, error) {
	return proxy.Service.{{ .Name }}(ctx, request)
}
{{ end }}
{{ end }}
`))
