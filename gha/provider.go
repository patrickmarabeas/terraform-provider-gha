package gha

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type ResourceProvider struct {
	BaseURL        string
	Pem            string
	AppID          string
	InstallationID string
	Token          string
}

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", "https://api.github.com/"),
				Description: "The GitHub Base API URL.",
			},
			"pem": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_PEM", nil),
				Description: "The GitHub App PEM string.",
			},
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_ID", nil),
				Description: "The GitHub App ID.",
			},
			"installation_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_INSTALLATION_ID", nil),
				Description: "The GitHub App installation instance ID.",
			},
		},
		ResourcesMap: nil,

		DataSourcesMap: map[string]*schema.Resource{
			"gha_token": dataSourceGhaToken(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := ResourceProvider{
			BaseURL:        d.Get("base_url").(string),
			Pem:            d.Get("pem").(string),
			AppID:          d.Get("app_id").(string),
			InstallationID: d.Get("installation_id").(string),
		}

		return config, nil
	}
}
