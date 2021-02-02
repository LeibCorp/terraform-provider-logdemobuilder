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
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"log": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	log := d.Get("log").(string)
	instance := d.Get("instance").(string)
	m := Message{log,instance}
	b, err := json.Marshal(m)
	if err == nil {
		resp, err := http.Post("https://daidemos.com/log", "application/json", bytes.NewBuffer(b))
		if err != nil {}
		defer resp.Body.Close()
	}
	return diags
}

