---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_storage_plan"
sidebar_current: "docs-alicloud-resource-polardb-_storage-plan"
description: |-
Provides a collection of PolarDB Storage Packages according to the specified filters.
---

# alicloud\_polardb\_storage_plan

_storage_plan

-> **NOTE:** Available in v1.140.0+.

## Example Usage

```terraform
resource "alicloud_polardb_storage_plan" "default" {
  storage_type = "Mainland"
  storage_class       = "50"
  period = 1
}
```

## Argument Reference

The following arguments are supported:

* `storage_type` - Specification of storage package type.Valid values are `Overseas`, `Mainland`, `Financial`.
* `storage_class` - Specification of storage package.Valid values are `50`, `100`, `200`, `300`, `500`, `1000`, `2000`, `3000`, `5000`, `10000`, `15000`, `20000`, `25000`, `30000`, `50000`, `100000`, `200000`.
* `period` - The duration that you will buy storage package (in month). Valid values: [1~9], 12, 24, 36, 60.
  -> **NOTE:** The argument `period` is only used to create Subscription storage package. Once effect, it will not be modified that means running `terraform apply` will not effect the resource.
## Attributes Reference

The following attributes are exported:

* `id` - The ID of the PolarDB Storage Package.
* `prod_code` - Product code.
* `ali_uid` - AliUid.
* `commodity_code` - Commodity code.
* `template_name` - Resource package type.
* `storage_type` - Specification of storage package type.Valid values are `Overseas`, `Mainland`, `Financial`.
* `storage_class` - Specification of storage package.Valid values are `50`, `100`, `200`, `300`, `500`, `1000`, `2000`, `3000`, `5000`, `10000`, `15000`, `20000`, `25000`, `30000`, `50000`, `100000`, `200000`.
* `other_property` - Specification of storage package remark.
* `status` - Status of the Storage Package.Valid values are `valid`, `invalid`.
* `start_times` - Resource start times.
* `end_times` - Resource end times.
* `purchase_times` - Resource purchase times.
* `init_capacity_view_value` - Initial capacity.
* `init_capa_city_view_unit` - Initial capacity variable unit.
* `period_capacity_view_value` - Resource cycle capacity.
* `period_capa_city_view_unit` - Resource cycle capacity unit.
* `period_time` - Cycle duration of the resource.

## Import

PolarDB storage package can be imported using the id, e.g.

```
$ terraform import alicloud_polardb_storage_plan.example "POLARDB-cn-123456"
```
