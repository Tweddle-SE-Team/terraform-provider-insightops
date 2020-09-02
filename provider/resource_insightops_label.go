package provider

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightOpsLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightOpsLabelCreate,
		Read:   resourceInsightOpsLabelRead,
		Delete: resourceInsightOpsLabelDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"sn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reserved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInsightOpsLabelCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	label := getInsightOpsLabelFromData(data)
	if err := client.PostLabel(label); err != nil {
		return err
	}
	if err := setInsightOpsLabelData(data, label); err != nil {
		return err
	}
	return resourceInsightOpsLabelRead(data, meta)
}

func resourceInsightOpsLabelImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightOpsLabelRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	label, err := client.GetLabel(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightOpsLabelData(data, label); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLabelUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	label := getInsightOpsLabelFromData(data)
	if err := client.PutLabel(label); err != nil {
		return err
	}
	if err := setInsightOpsLabelData(data, label); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLabelDelete(data *schema.ResourceData, meta interface{}) error {
	labelId := data.Id()
	client := meta.(*insightops.InsightOpsClient)
	if err := client.DeleteLabel(labelId); err != nil {
		return err
	}
	return nil
}

func getInsightOpsLabelFromData(data *schema.ResourceData) *insightops.Label {
	return &insightops.Label{
		Id:       data.Id(),
		Color:    data.Get("color").(string),
		SN:       data.Get("sn").(int),
		Reserved: data.Get("reserved").(bool),
	}
}

func setInsightOpsLabelData(data *schema.ResourceData, label *insightops.Label) error {
	data.SetId(label.Id)
	data.Set("color", label.Color)
	data.Set("sn", label.SN)
	data.Set("reserved", label.Reserved)
	return nil
}
