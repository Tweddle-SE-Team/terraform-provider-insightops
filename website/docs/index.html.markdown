---
layout: "insightops"
page_title: "Provider: InsightOps"
sidebar_current: "docs-insightops-index"
description: |-
  The InsightOps provider is used to interact with the resources supported by Rapid7 InsightOps. The provider needs to be configured with the proper credentials before it can be used.
---

# InsightOps Provider

The InsightOps provider is used to interact with the resources supported by Rapid7 InsightOps. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "insightops" {
  api_key = "insight-ops-api-key"
  region  = "us"
}

resource "insightops_logset" "logset" {
  name        = "logset"
  description = "test logset"
}
```

## Argument Reference

The following arguments are supported:

* `endpoint` - (Optional) InsightOps API endpoint. Can be sourced from `INSIGHTOPS_ENDPOINT`.
* `api_key` - (Required) InsightOps API Key. Can be sourced from `INSIGHTOPS_API_KEY`.
* `region` - (Required) InsightOps API region. Can be sourced from `INSIGHTOPS_REGION`.
