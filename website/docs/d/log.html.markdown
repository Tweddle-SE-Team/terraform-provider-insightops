---
layout: "insightops"
page_title: "InsightOps: insightops_log"
sidebar_current: "docs-insightops-data-source-log"
description: |-
  Gets log info.
---

# insightops_log

This data source provides a mechanism for retrieving an information about Rapid7 InsightOps log.

## Example Usage

```hcl
data "insightops_log" "log" {
    name   = "test"
    logset = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps log name.
* `logset` - (Required) Rapid7 InsightOps logset name.

## Attribute Reference

* `id` - Log Unique Identifier
* `token` - Log token
