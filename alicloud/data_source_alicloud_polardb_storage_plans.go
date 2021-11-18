package alicloud

import (
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAlicloudPolarDBStoragePlans() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudStoragePlansPolarDBRead,

		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_plans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prod_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ali_uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"commodity_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_times": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_times": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purchase_times": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"init_capacity_view_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"init_capa_city_view_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"period_capacity_view_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"period_capa_city_view_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"period_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAlicloudStoragePlansPolarDBRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "DescribeStoragePlan"
	request := map[string]interface{}{
		"PageSize":   PageSizeLarge,
		"PageNumber": 1,
	}

	var response map[string]interface{}
	conn, err := client.NewPolarDBClient()
	if err != nil {
		return WrapError(err)
	}

	var objects []map[string]interface{}

	for {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-08-01"), StringPointer("AK"), nil, request, &runtime)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_polardb_storage_plans", action, AlibabaCloudSdkGoERROR)
		}
		resp, err := jsonpath.Get("$.Items", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, d.Id(), response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			objects = append(objects, item)
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	var ids []string
	var s []map[string]interface{}

	for _, item := range objects {
		mapping := map[string]interface{}{
			"id":                         fmt.Sprint(item["InstanceId"]),
			"prod_code":                  fmt.Sprint(item["ProdCode"]),
			"ali_uid":                    fmt.Sprint(item["AliUid"]),
			"commodity_code":             fmt.Sprint(item["CommodityCode"]),
			"template_name":              fmt.Sprint(item["TemplateName"]),
			"storage_type":               fmt.Sprint(item["StorageType"]),
			"status":                     fmt.Sprint(item["Status"]),
			"start_times":                fmt.Sprint(item["StartTimes"]),
			"end_times":                  fmt.Sprint(item["EndTimes"]),
			"purchase_times":             fmt.Sprint(item["PurchaseTimes"]),
			"init_capacity_view_value":   fmt.Sprint(item["InitCapacityViewValue"]),
			"init_capa_city_view_unit":   fmt.Sprint(item["InitCapaCityViewUnit"]),
			"period_capacity_view_value": fmt.Sprint(item["PeriodCapacityViewValue"]),
			"period_capa_city_view_unit": fmt.Sprint(item["PeriodCapaCityViewUnit"]),
			"period_time":                fmt.Sprint(item["PeriodTime"]),
		}
		ids = append(ids, fmt.Sprint(item["InstanceName"]))
		s = append(s, mapping)
	}
	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}
	if err := d.Set("storage_plans", s); err != nil {
		return WrapError(err)
	}

	// create a json file in current directory and write data source to it
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}
