package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var apiKey = os.Getenv("INSIGHTOPS_API_KEY")
var region = os.Getenv("INSIGHTOPS_REGION")

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"insightops": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("INSIGHTOPS_API_KEY") == "" {
		t.Fatalf("err: INSIGHTOPS_API_KEY env variable is mandatory to run acceptance tests")
	}
	if os.Getenv("INSIGHTOPS_REGION") == "" {
		t.Fatalf("err: INSIGHTOPS_REGION env variable is mandatory to run acceptance tests")
	}
}
