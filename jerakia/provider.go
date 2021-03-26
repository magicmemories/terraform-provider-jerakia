package jerakia

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//"github.com/jerakia/go-jerakia"
	"github.com/magicmemories/go-jerakia"
)

// Provider returns a schema.Provider for Jerakia
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Description: "The URL to the Jerakia service. This can also be set with the `JERAKIA_URL` environment variable.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JERAKIA_URL", ""),
			},

			"api_token": &schema.Schema{
				Description: "The token to authenticate to Jerakia with. This can also be set with the `JERAKIA_TOKEN` environment variable.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JERAKIA_TOKEN", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"jerakia_lookup": dataSourceLookup(),
		},

		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	api_url := d.Get("api_url").(string)
	api_token := d.Get("api_token").(string)

	var diags diag.Diagnostics

	config := jerakia.ClientConfig{
		URL:   api_url,
		Token: api_token,
	}

	client := jerakia.NewClient(http.DefaultClient, config)

	return client, diags
}
