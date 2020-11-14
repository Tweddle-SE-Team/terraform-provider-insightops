package provider

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// InsightOpsProvider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHTOPS_API_KEY", nil),
				Description: "API key (Read/Write) to be able to interact with InsightOps REST API",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHTOPS_REGION", nil),
				Description: "Region for InsightOps REST API",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHTOPS_ENDPOINT", nil),
				Description: "Custom InsightOps REST API Endpoint",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"insightops_label":  dataSourceInsightOpsLabel(),
			"insightops_log":    dataSourceInsightOpsLog(),
			"insightops_logset": dataSourceInsightOpsLogset(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"insightops_tag":    resourceInsightOpsTag(),
			"insightops_logset": resourceInsightOpsLogset(),
			"insightops_log":    resourceInsightOpsLog(),
			"insightops_label":  resourceInsightOpsLabel(),
			"insightops_action": resourceInsightOpsAction(),
			"insightops_target": resourceInsightOpsTarget(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	apiKey := data.Get("api_key").(string)
	region := data.Get("region").(string)
	client, err := insightops.NewInsightOpsClient(apiKey, region)
	if err != nil {
		return nil, err
	}
	if endpoint, ok := data.GetOk("endpoint"); ok {
		client.InsightOpsUrl = endpoint.(string)
	}
	return client, nil
}
