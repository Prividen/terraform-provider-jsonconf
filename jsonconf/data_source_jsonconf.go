package jsonconf

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"strconv"
	"time"
)

func dataSourceJsonConf() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJsonConfRead,
		Schema: map[string]*schema.Schema{
			"nodes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"codename": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_vm": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceJsonConfRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics


	providerConfig := m.(*Config)
	file := providerConfig.File

	jsonFile, err := os.Open(file)
	if err != nil {
		return diag.FromErr(err)
	}
	defer jsonFile.Close()

	jsonConf := make([]map[string]interface{}, 0)
	err = json.NewDecoder(jsonFile).Decode(&jsonConf)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("nodes", jsonConf); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}