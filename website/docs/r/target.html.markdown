---
layout: "insightops"
page_title: "InsightOps: insightops_target"
sidebar_current: "docs-insightops-resource-target"
description: |-
  Creates InsightOps Target
---

# insightops_target

This resource provides a mechanism to create Rapid7 InsightOps target.

## Example Usage

```hcl
resource insightops_target pagerduty {
  name = "pgdt"
  pagerduty_service_key = "asdasdasdasd"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps target name.
* `pagerduty_service_key` - (Optional) Pagetduty service key
* `webhook_url` - (Optional) Custom webhook url
* `slack_webhook` - (Optional) Slack incoming webhook
* `log_context` - (Optional) Enable log context
* `action_ids` - (Optional) Action ids
* `log_link` - (Optional) Include log link

## Attribute Reference

* `id` - Rapid7 Target Unique Identifier