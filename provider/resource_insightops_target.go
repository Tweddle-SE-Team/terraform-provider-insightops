package provider

import (
	"fmt"
	"github.com/AndrewChubatiuk/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightOpsTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightOpsTargetCreate,
		Read:   resourceInsightOpsTargetRead,
		Delete: resourceInsightOpsTargetDelete,
		Update: resourceInsightOpsTargetUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightOpsTargetImport,
		},
		Schema: map[string]*schema.Schema{
			"pagerduty_service_key": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"webhook_url", "slack_webhook"},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"webhook_url": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"slack_webhook": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"pagerduty_service_key", "slack_webhook"},
			},
			"log_context": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"log_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInsightOpsTargetCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	target := getInsightOpsTargetFromData(data)
	if err := client.PostTarget(target); err != nil {
		return err
	}
	if err := setInsightOpsTargetData(data, target); err != nil {
		return err
	}
	return resourceInsightOpsTargetRead(data, meta)
}

func resourceInsightOpsTargetImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightOpsTargetRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	target, err := client.GetTarget(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightOpsTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	target := getInsightOpsTargetFromData(data)
	if err := client.PutTarget(target); err != nil {
		return err
	}
	if err := setInsightOpsTargetData(data, target); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsTargetDelete(data *schema.ResourceData, meta interface{}) error {
	targetId := data.Id()
	client := meta.(*insightops.InsightOpsClient)
	if err := client.DeleteTarget(targetId); err != nil {
		return err
	}
	return nil
}

func getInsightOpsTargetFromData(data *schema.ResourceData) *insightops.Target {
	target := insightops.Target{
		Id:   data.Id(),
		Name: data.Get("name").(string),
		AlertContentSet: &insightops.TargetAlertContentSet{
			LogLink: insightops.StringBool(data.Get("log_link").(bool)),
			Context: insightops.StringBool(data.Get("log_context").(bool)),
		},
	}
	if attr, ok := data.GetOk("pagerduty_service_key"); ok {
		target.Type = "Pagerduty"
		target.ParameterSet = &insightops.TargetParameterSet{
			ServiceKey: attr.(string),
		}
	} else if attr, ok := data.GetOk("webhook_url"); ok {
		target.Type = "Webhook"
		target.ParameterSet = &insightops.TargetParameterSet{
			Url: attr.(string),
		}
	} else if attr, ok := data.GetOk("slack_webhook"); ok {
		target.Type = "Slack"
		target.ParameterSet = &insightops.TargetParameterSet{
			Url: attr.(string),
		}
	}
	return &target
}

func setInsightOpsTargetData(data *schema.ResourceData, target *insightops.Target) error {
	data.SetId(target.Id)
	data.Set("log_link", target.AlertContentSet.LogLink)
	data.Set("log_context", target.AlertContentSet.Context)
	data.Set("name", target.Name)
	switch target.Type {
	case "Webhook":
		data.Set("webhook_url", target.ParameterSet.Url)
	case "Slack":
		data.Set("slack_webhook", target.ParameterSet.Url)
	case "Pagerduty":
		data.Set("pagerduty_service_key", target.ParameterSet.ServiceKey)
	default:
		return fmt.Errorf("%s target type is not supported", target.Type)
	}
	return nil
}
