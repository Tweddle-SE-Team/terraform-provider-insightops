package provider

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceInsightOpsLog() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logset": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceLogRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	name := data.Get("name").(string)
	logset := data.Get("logset").(string)
	token, id, err := client.GetLogToken(logset, name)
	if err != nil {
		return err
	}
	data.SetId(id)
	data.Set("token", token)
	return nil
}
