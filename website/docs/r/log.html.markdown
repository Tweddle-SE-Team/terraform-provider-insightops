---
layout: "insightops"
page_title: "InsightOps: insightops_log"
sidebar_current: "docs-insightops-resource-log"
description: |-
  Creates InsightOps Log
---

# insightops_log

This resource provides a mechanism to create Rapid7 InsightOps log.

## Example Usage

```hcl
data "insightops_logset" "logset" {
    name = "test"
}

resource "insightops_log" "log" {
  name           = "My super log"
  source_type    = "token"
  logset_ids     = [data.insightops_logset.logset.id]
  agent_filename = "/var/log/anaconda.log"
  agent_follow   = "true"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps label name.
* `source_type` - (Required) Rapid7 InsightOps label color.
* `logset_ids` - (Required) Rapid7 InsightOps label color.
* `agent_filename` - (Required) Rapid7 InsightOps label color.
* `agent_follow` - (Required) Rapid7 InsightOps label color.
* `token_seed` - (Optional) Log token seed

## Attribute Reference

* `id` - Rapid7 Log Unique Identifier
* `tokens` - Log tokens 