---
layout: "insightops"
page_title: "InsightOps: insightops_action"
sidebar_current: "docs-insightops-resource-action"
description: |-
  Creates InsightOps Action
---

# insightops_action

This resource provides a mechanism to create Rapid7 InsightOps action.

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

resource "insightops_action" "alert1" {
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
* `type` - (Optional) Action type
* `enabled` - (Optional) Action status
* `min_matches_count` - (Optional) Minimal matches action count
* `min_matches_period` - (Optional) Minimal matches action period
* `min_report_count`- (Optional) Minimal report action count
* `min_report_period` - (Optional) Minimal report action period
* `target_ids` - (Optional) Rapid7 Target IDs

## Attribute Reference

* `id` - Rapid7 Action Unique Identifier