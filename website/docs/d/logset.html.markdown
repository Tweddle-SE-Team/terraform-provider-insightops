---
layout: "insightops"
page_title: "InsightOps: insightops_logset"
sidebar_current: "docs-insightops-data-source-logset"
description: |-
  Gets logset info.
---

# insightops_logset

This data source provides a mechanism for retrieving an information about Rapid7 InsightOps logset.

## Example Usage

```hcl
data "insightops_logset" "logset" {
    name = "test"
}

resource "insightops_log" "log" {
  name           = "My super log"
  source_type    = "token"
  logset_ids     = [insightops_logset.logset.id]
  agent_filename = "/var/log/anaconda.log"
  agent_follow   = "true"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps logset name.

## Attribute Reference

* `description` - Logset description.
* `id` - Logset Unique Identifier
