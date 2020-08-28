package provider

import (
	"fmt"
	"github.com/AndrewChubatiuk/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

const tagsResourceName = "insightops_tag"
const tagsResourceId = "acceptance_tag"

var tagsResourceStateId = fmt.Sprintf("%s.%s", tagsResourceName, tagsResourceId)

var testTagsCreateConfig string
var testTagsUpdatedConfig string

var createTagName = "My App Failures"
var createTagPatterns = "[error]"
var createTagActionParamSetDescription = "Log Error"

var updatedTagName = "My Update App Failures"
var updatedTagPatterns = "[error] AND failure and 502"
var updatedTagActionParamSetDescription = "Log Error description updated"

var sourceId = os.Getenv("SOURCE_ID")

func init() {

	configTemplate := `
	provider insightops {
	  api_key = "%s"
	  region  = "%s"
	}

	resource insightops_logset acceptance_logset {
	  name = "Sample LogSet for tag acc tests"
	  description = "some description goes here"
	}

	resource insightops_log acceptance_log {
	  name = "Sample Log for tag acc tests"
	  source_type = "token"
	  token_seed = ""
	  structures = []
	  logsets_info = ["${insightops_logset.acceptance_logset.id}"]
	  user_data = {
	   	agent_filename = ""
			agent_follow = false
	  }
	}

	resource %s %s {
	  name = "%s"
	  type = "Alert"
	  patterns = ["%s"]
	  source_ids = ["${insightops_log.acceptance_log.id}"]
	  label_ids = []
	  action_ids = []
	}`

	testTagsCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, tagsResourceName, tagsResourceId, createTagName, createTagPatterns, createTagActionParamSetDescription)
	testTagsUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, region, tagsResourceName, tagsResourceId, updatedTagName, updatedTagPatterns, updatedTagActionParamSetDescription)
}

func tagExists() checkExists {
	return func(client insightops.InsightOpsClient, id string) error {
		_, err := client.GetTag(id)
		return err
	}
}

func TestAccLogentriesTags_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep{
			{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
				),
			},
		},
	})
}

func TestAccLogentriesTags_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(tagsResourceStateId, tagExists()),
		Steps: []resource.TestStep{
			{
				Config: testTagsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", createTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", createTagPatterns),
				),
			},
			{
				Config: testTagsUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tagsResourceStateId, "name", updatedTagName),
					resource.TestCheckResourceAttr(tagsResourceStateId, "type", "Alert"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.#", "1"),
					resource.TestCheckResourceAttr(tagsResourceStateId, "patterns.0", updatedTagPatterns),
				),
			},
		},
	})
}
