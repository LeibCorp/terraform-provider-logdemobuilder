package provider

import (
	"context"
	"net/http"
	"encoding/json"
	"bytes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Message struct {
    Log string
    Instance string
    Ip string
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"log": {
				Type:     schema.TypeString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"instance": {
				Type:     schema.TypeString,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	m := Message{d.Get("log").(string),d.Get("instance").(string),d.Get("ip").(string)}
	b, err := json.Marshal(m)
	if err == nil {
		resp, err := http.Post("https://daitest.leibcorp.com/log", "application/json", bytes.NewBuffer(b))
		if err != nil {}
		defer resp.Body.Close()
	}
	return diags
}

