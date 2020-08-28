---
layout: "insightops"
page_title: "InsightOps: insightops_label"
sidebar_current: "docs-insightops-data-source-label"
description: |-
  Gets label info.
---

# insightops_label

This data source provides a mechanism for retrieving an information about Rapid7 InsightOps label.

## Example Usage

```hcl
data "insightops_label" "label" {
  name  = "Critical"
  color = "e0e000"
}

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

resource "insightops_action" "alert" {
  type               = "Alert"
  enabled            = true
  min_matches_count  = 1
  min_matches_period = "Hour"
  min_report_count   = 1
  min_report_period  = "Hour"
  target_ids         = [insightops_target.pagerduty.id]
}

resource "insightops_target" "pagerduty" {
  name = "pgdt"
  pagerduty_service_key = "asdasdasdasd"
}

resource "insightops_tag" "tag" {
  name       = "My App Failures"
  type       = "Alert"
  patterns   = ["[error]"]
  source_ids = [insightops_log.log.id]
  label_ids  = [data.insightops_label.label.id]
  action_ids = [insightops_action.alert.id]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps label name.
* `color` - (Optional) Rapid7 InsightOps label color.

## Attribute Reference

* `id` - Logset Unique Identifier