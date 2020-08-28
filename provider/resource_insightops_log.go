package provider

import (
	"github.com/AndrewChubatiuk/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInsightOpsLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceInsightOpsLogCreate,
		Read:   resourceInsightOpsLogRead,
		Delete: resourceInsightOpsLogDelete,
		Update: resourceInsightOpsLogUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logset_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_seed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tokens": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"structures": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"agent_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_follow": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInsightOpsLogCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	log := getInsightOpsLogFromData(data)
	if err := client.PostLog(log); err != nil {
		return err
	}
	if err := setInsightOpsLogData(data, log); err != nil {
		return err
	}
	return resourceInsightOpsLogRead(data, meta)
}

func resourceInsightOpsLogImport(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{data}, nil
}

func resourceInsightOpsLogRead(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	log, err := client.GetLog(data.Id())
	if err != nil {
		return nil
	}
	if err = setInsightOpsLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLogUpdate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*insightops.InsightOpsClient)
	log := getInsightOpsLogFromData(data)
	if err := client.PutLog(log); err != nil {
		return err
	}
	if err := setInsightOpsLogData(data, log); err != nil {
		return err
	}
	return nil
}

func resourceInsightOpsLogDelete(data *schema.ResourceData, meta interface{}) error {
	logId := data.Id()
	client := meta.(*insightops.InsightOpsClient)
	if err := client.DeleteLog(logId); err != nil {
		return err
	}
	return nil
}

func getInsightOpsLogFromData(data *schema.ResourceData) *insightops.Log {
	var structures []string
	var logsetIds []*insightops.Info
	var tokens []string
	if v, ok := data.GetOk("structures"); ok {
		for _, structure := range v.(*schema.Set).List() {
			structures = append(structures, structure.(string))
		}
	}
	if v, ok := data.GetOk("tokens"); ok {
		for _, token := range v.(*schema.Set).List() {
			tokens = append(tokens, token.(string))
		}
	}
	if v, ok := data.GetOk("logset_ids"); ok {
		for _, logsetId := range v.(*schema.Set).List() {
			logsetIds = append(logsetIds, &insightops.Info{Id: logsetId.(string)})
		}
	}
	return &insightops.Log{
		Id:          data.Id(),
		SourceType:  data.Get("source_type").(string),
		Name:        data.Get("name").(string),
		TokenSeed:   data.Get("token_seed").(string),
		Tokens:      tokens,
		Structures:  structures,
		LogsetsInfo: logsetIds,
		UserData: &insightops.LogUserData{
			AgentFileName: data.Get("agent_filename").(string),
			AgentFollow:   insightops.StringBool(data.Get("agent_follow").(bool)),
		},
	}
}

func setInsightOpsLogData(data *schema.ResourceData, log *insightops.Log) error {
	var logsets []string
	for _, logset := range log.LogsetsInfo {
		logsets = append(logsets, logset.Id)
	}
	data.Set("agent_follow", log.UserData.AgentFollow)
	data.Set("agent_filename", log.UserData.AgentFileName)
	data.SetId(log.Id)
	data.Set("name", log.Name)
	data.Set("source_type", log.SourceType)
	data.Set("token_seed", log.TokenSeed)
	data.Set("tokens", log.Tokens)
	data.Set("structures", log.Structures)
	data.Set("logset_ids", logsets)
	return nil
}
