package newthings

import (
	"context"
	_ "embed"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure NewBrokeProvider satisfies various provider interfaces.
var _ provider.Provider = &NewBrokeProvider{}

//go:embed moby.dick.txt
var mobyDickText string

// NewBrokeProvider defines the provider implementation.
type NewBrokeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NewBrokeProviderModel describes the provider data model.
type NewBrokeProviderModel struct {
}

func (p *NewBrokeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "broke"
	resp.Version = p.version
}

func (p *NewBrokeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *NewBrokeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data NewBrokeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *NewBrokeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewMobyDickResource,
		NewDriftAgainResource,
	}
}

func (p *NewBrokeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) provider.Provider {
	return &NewBrokeProvider{
		version: version,
	}
}
