---
layout: "insightops"
page_title: "InsightOps: insightops_label"
sidebar_current: "docs-insightops-resource-label"
description: |-
  Creates InsightOps Label
---

# insightops_label

This resource provides a mechanism to create Rapid7 InsightOps label.

## Example Usage

```hcl
resource "insightops_label" "label" {
  name  = "Label"
  color = "ff0000"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Rapid7 InsightOps label name.
* `color` - (Required) Rapid7 InsightOps label color.

## Attribute Reference

* `id` - Rapid7 Label Unique Identifier