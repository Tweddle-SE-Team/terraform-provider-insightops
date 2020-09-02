package provider

import (
	"fmt"
	"github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const logsResourceName = "insightops_log"
const logsResourceId = "acceptance_log"

var logResourceStateId = fmt.Sprintf("%s.%s", logsResourceName, logsResourceId)

var testLogCreateConfig string
var testLogUpdatedConfig string

var createLogName = "My New Awesome Log"
var createLogAgentFileName = "/var/log/anaconda.log"
var createLogAgentFollow = "false"

var updatedLogName = "My Updated Awesome Log"
var updatedLogAgentFileName = "/var/log/snake.log"
var updatedLogAgentFollow = "true"

func init() {

	configTemplate := `
	provider "insightops" {
	  api_key = "%s"
	  region  = "%s"
	}

	resource insightops_logset acceptance_logset {
	  name = "Acceptance Log Set for log acc tests"
	  description = "some description goes here"
	}

	resource "%s" "%s" {
	  	name = "%s"
	  	source_type = "token"
	  	token_seed = ""
	  	structures = []
	  	logsets_info = ["${insightops_logset.acceptance_logset.id}"]
	  	agent_filename = "%s"
		agent_follow   = "%s"
	}`

	testLogCreateConfig = fmt.Sprintf(configTemplate, apiKey, region, logsResourceName, logsResourceId, createLogName, createLogAgentFileName, createLogAgentFollow)
	testLogUpdatedConfig = fmt.Sprintf(configTemplate, apiKey, region, logsResourceName, logsResourceId, updatedLogName, updatedLogAgentFileName, updatedLogAgentFollow)
}

func logExists() checkExists {
	return func(client insightops.InsightOpsClient, id string) error {
		_, err := client.GetLog(id)
		return err
	}
}

func TestAccInsightOpsLog_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep{
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_follow", createLogAgentFollow),
				),
			},
		},
	})
}

func TestAccInsightOpsLog_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkDestroy(logResourceStateId, logExists()),
		Steps: []resource.TestStep{
			{
				Config: testLogCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", createLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_filename", createLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_follow", createLogAgentFollow),
				),
			},
			{
				Config: testLogUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logResourceStateId, "name", updatedLogName),
					resource.TestCheckResourceAttr(logResourceStateId, "source_type", "token"),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_filename", updatedLogAgentFileName),
					resource.TestCheckResourceAttr(logResourceStateId, "agent_follow", updatedLogAgentFollow),
				),
			},
		},
	})
}
