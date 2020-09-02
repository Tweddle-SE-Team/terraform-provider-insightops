package provider

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightOpsLogset() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightOpsLogsetCreate,
		Read:   resourceInsightOpsLogsetRead,
		Delete: resourceInsightOpsLogsetDelete,
		Update: resourceInsightOpsLogsetUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightOpsLogsetImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceInsightOpsLogsetCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	logset := getInsightOpsLogsetFromData(data)
	if err := client.PostLogset(logset); err != nil {
		return err
	}
	if err := setInsightOpsLogsetData(data, logset); err != nil {
		return err
	}
	return resourceInsightOpsLogsetRead(data, meta)
}

func resourceInsightOpsLogsetImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightOpsLogsetRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	logset, err := client.GetLogset(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightOpsLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLogsetUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	logset := getInsightOpsLogsetFromData(data)
	logsetInfo, err := client.GetLogset(data.Id())
	if err != nil {
		return nil
	}
	logset.LogsInfo = logsetInfo.LogsInfo
	if err := client.PutLogset(logset); err != nil {
		return err
	}
	if err := setInsightOpsLogsetData(data, logset); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLogsetDelete(data *schema.ResourceData, meta interface{}) error {
	logsetId := data.Id()
	client := meta.(*insightops.InsightOpsClient)
	if err := client.DeleteLogset(logsetId); err != nil {
		return err
	}
	return nil
}

func getInsightOpsLogsetFromData(data *schema.ResourceData) *insightops.Logset {
	return &insightops.Logset{
		Id:          data.Id(),
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
	}
}

func setInsightOpsLogsetData(data *schema.ResourceData, logset *insightops.Logset) error {
	data.SetId(logset.Id)
	data.Set("name", logset.Name)
	data.Set("description", logset.Description)
	return nil
}
