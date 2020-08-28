---
layout: "insightops"
page_title: "InsightOps: insightops_tag"
sidebar_current: "docs-insightops-resource-tag"
description: |-
  Creates InsightOps Tag
---

# insightops_tag

This resource provides a mechanism to create Rapid7 InsightOps tag.

## Example Usage

```hcl
data "insightops_label" "label" {
  name  = "Critical"
  color = "e0e000"
}

resource "insightops_tag" "tag" {
  name       = "App Failures"
  type       = "Alert"
  patterns   = ["[error]"]
  source_ids = ["5a1288ab-561a-4f93-1111-6a38c6d8TEST"]
  label_ids  = [data.insightops_label.label.id]
  action_ids = [insightops_action.alert.id]
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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps tag name.
* `type` - (Optional) Rapid7 InsightOps tag type
* `patterns` - (Optional) Tag pattern
* `source_ids` - (Optional) Log ids
* `label_ids` - (Optional) Label ids
* `action_ids` - (Optional) Action ids

## Attribute Reference

* `id` - Rapid7 Tag Unique Identifier