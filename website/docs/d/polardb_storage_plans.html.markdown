---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_storage_plans"
sidebar_current: "docs-alicloud-datasource-polardb-storage-plans"
description: |-
    Provides a collection of PolarDB Storage Packages according to the specified filters.
---

# alicloud\_polardb\_storage_plans

The `alicloud_polardb_storage_plans` data source provides a collection of PolarDB Storage Packages available in Alibaba Cloud account.
Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.

-> **NOTE:** Available in v1.140.0+.

## Argument Reference

The following arguments are supported:

* `ids` - (Optional) A list of PolarDB Storage Package IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of Storage Package IDs.
* `describeStoragePlanResponses` - A list of PolarDB Storage Packages. Each element contains the following attributes:
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
  