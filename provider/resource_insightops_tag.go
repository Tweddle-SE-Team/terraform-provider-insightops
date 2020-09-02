package provider

import (
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightOpsTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightOpsTagCreate,
		Read:   resourceInsightOpsTagRead,
		Delete: resourceInsightOpsTagDelete,
		Update: resourceInsightOpsTagUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceInsightOpsTagImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"action_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"patterns": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"label_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceInsightOpsTagCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	tag := getInsightOpsTagFromData(data)
	if err := client.PostTag(tag); err != nil {
		return err
	}
	if err := setInsightOpsTagData(data, tag); err != nil {
		return err
	}
	return resourceInsightOpsTagRead(data, meta)
}

func resourceInsightOpsTagImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightOpsTagRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	tag, err := client.GetTag(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightOpsTagData(data, tag); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsTagUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	tag := getInsightOpsTagFromData(data)
	if err := client.PutTag(tag); err != nil {
		return err
	}
	if err := setInsightOpsTagData(data, tag); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsTagDelete(data *schema.ResourceData, meta interface{}) error {
	tagId := data.Id()
	client := meta.(*insightops.InsightOpsClient)
	if err := client.DeleteTag(tagId); err != nil {
		return err
	}
	return nil
}

func getInsightOpsTagFromData(data *schema.ResourceData) *insightops.Tag {
	var patterns []string
	if v, ok := data.GetOk("patterns"); ok {
		for _, pattern := range v.(*schema.Set).List() {
			patterns = append(patterns, pattern.(string))
		}
	}
	var actions []*insightops.Action
	for _, id := range data.Get("action_ids").(*schema.Set).List() {
		actions = append(actions, &insightops.Action{Id: id.(string)})
	}
	var sources []*insightops.Source
	for _, id := range data.Get("source_ids").(*schema.Set).List() {
		sources = append(sources, &insightops.Source{Id: id.(string)})
	}
	var labels []*insightops.Label
	for _, id := range data.Get("label_ids").(*schema.Set).List() {
		labels = append(labels, &insightops.Label{Id: id.(string)})
	}
	return &insightops.Tag{
		Id:       data.Id(),
		Type:     data.Get("type").(string),
		Name:     data.Get("name").(string),
		Patterns: patterns,
		Sources:  sources,
		Actions:  actions,
		Labels:   labels,
		UserData: map[string]string{"product_type": "OPS"},
	}
}

func setInsightOpsTagData(data *schema.ResourceData, tag *insightops.Tag) error {
	var labels []string
	for _, label := range tag.Labels {
		labels = append(labels, label.Id)
	}
	var sources []string
	for _, source := range tag.Sources {
		sources = append(sources, source.Id)
	}
	var actions []string
	for _, action := range tag.Actions {
		actions = append(actions, action.Id)
	}
	data.SetId(tag.Id)
	data.Set("patterns", tag.Patterns)
	data.Set("name", tag.Name)
	data.Set("type", tag.Type)
	data.Set("patterns", tag.Patterns)
	data.Set("label_ids", labels)
	data.Set("source_ids", sources)
	data.Set("action_ids", actions)
	return nil
}
