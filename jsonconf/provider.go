package jsonconf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
type Config struct {
	File  string
}
// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	"conf.json",
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"jsonconf_nodes": dataSourceJsonConf(),
		},
		ConfigureContextFunc: providerConfigure(),
	}
}

func providerConfigure() schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics
		config := &Config{}
		config.File = d.Get("file").(string)
		return config, diags
	}
}