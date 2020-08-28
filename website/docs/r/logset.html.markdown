---
layout: "insightops"
page_title: "InsightOps: insightops_logset"
sidebar_current: "docs-insightops-resource-logset"
description: |-
  Creates InsightOps Logset
---

# insightops_logset

This resource provides a mechanism to create Rapid7 InsightOps logset.

## Example Usage

```hcl
resource "insightops_logset" "logset" {
  name        = "Log Set"
  description = "Description about my log set"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps logset name.
* `description` - (Optional) Rapid7 InsightOps logset description.

## Attribute Reference

* `id` - Rapid7 Logset Unique Identifier