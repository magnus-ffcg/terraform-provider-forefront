package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ForefrontProvider satisfies various provider interfaces.
var _ provider.Provider = &ForefrontProvider{}

// ForefrontProvider defines the provider implementation.
type ForefrontProvider struct {
	// version is set to the provider version on release, "dev" when in the
	// testing environment, and empty when compiled from source.
	version string
}

// ForefrontProviderModel describes the provider data model.
type ForefrontProviderModel struct {
	BaseURL types.String `tfsdk:"base_url"`
	APIKey  types.String `tfsdk:"api_key"`
}

func (p *ForefrontProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "forefront"
	resp.Version = p.version
}

func (p *ForefrontProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				MarkdownDescription: "Base URL for the Forefront API. May also be provided via FOREFRONT_BASE_URL environment variable.",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API Key for Forefront authentication. May also be provided via FOREFRONT_API_KEY environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *ForefrontProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Example configuration logic.
	// You might read environment variables or the configuration block to initialize an API client here.
	
	// var data ForefrontProviderModel
	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// Example: initialize client
	// client := forefront.NewClient(data.APIKey.ValueString(), data.BaseURL.ValueString())
	
	// Example: pass client to resources/data sources
	// resp.DataSourceData = client
	// resp.ResourceData = client
}

func (p *ForefrontProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewLookupFunction,
		NewReplaceDeepFunction,
	}
}

func (p *ForefrontProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// TODO: Add resources here
	}
}

func (p *ForefrontProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// TODO: Add data sources here
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ForefrontProvider{
			version: version,
		}
	}
}
