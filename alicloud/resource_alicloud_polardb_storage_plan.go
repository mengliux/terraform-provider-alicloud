package alicloud

import (
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strconv"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudPolarDBStoragePlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudPolarDBStoragePlanCreate,
		Read:   resourceAlicloudPolarDBStoragePlanRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"storage_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Mainland", "Overseas"}, false),
				Required:     true,
			},
			"period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60}),
				Required:     true,
			},
			"storage_class": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"50", "100", "200", "300", "500", "1000", "2000", "3000", "5000", "10000", "15000", "20000", "25000", "30000", "50000", "100000", "200000"}, false),
				Required:     true,
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
			"other_property": {
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
	}
}

func resourceAlicloudPolarDBStoragePlanCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	conn, err := client.NewPolarDBClient()
	if err != nil {
		return WrapError(err)
	}
	action := "CreateStoragePlan"
	request, err := buildDBCreateStoragePlanRequest(d, meta)
	if err != nil {
		return WrapError(err)
	}
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
	d.SetId(response["DBInstanceId"].(string))
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}
	return resourceAlicloudPolarDBDatabaseRead(d, meta)
}

func resourceAlicloudPolarDBStoragePlanRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polarDBService := PolarDBService{client}
	storagePlan, err := polarDBService.DescribePolarDBStoragePlan(d)
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("id", storagePlan["InstanceName"])
	d.Set("ali_uid", storagePlan["AliUid"])
	d.Set("storage_type", storagePlan["StorageType"])
	d.Set("period", d.Get("period"))
	d.Set("storage_class", d.Get("storage_class"))
	d.Set("prod_code", storagePlan["ProdCode"])
	d.Set("commodity_code", storagePlan["CommodityCode"])
	d.Set("template_name", storagePlan["TemplateName"])
	d.Set("other_property", storagePlan["OtherProperty"])
	d.Set("status", storagePlan["Status"])
	d.Set("start_times", storagePlan["StartTimes"])
	d.Set("end_times", storagePlan["EndTimes"])
	d.Set("purchase_times", storagePlan["PurchaseTimes"])
	d.Set("init_capacity_view_value", storagePlan["InitCapacityViewValue"])
	d.Set("init_capa_city_view_unit", storagePlan["InitCapaCityViewUnit"])
	d.Set("period_capacity_view_value", storagePlan["PeriodCapacityViewValue"])
	d.Set("period_capa_city_view_unit", storagePlan["PeriodCapaCityViewUnit"])
	d.Set("period_time", storagePlan["PeriodTime"])

	return nil
}

func buildDBCreateStoragePlanRequest(d *schema.ResourceData, meta interface{}) (map[string]interface{}, error) {
	var request map[string]interface{}
	if storageClass, ok := d.GetOk("storage_class"); ok && Trim(storageClass.(string)) != "" {
		request["StorageClass"] = Trim(storageClass.(string))
	}
	if storageType, ok := d.GetOk("storage_type"); ok && Trim(storageType.(string)) != "" {
		request["StorageClass"] = Trim(storageType.(string))
	}
	// At present, API supports two charge options about 'Prepaid'.
	// 'Month': valid period ranges [1-9]; 'Year': valid period range [1,2,3,5]
	// This resource only supports to input Month period [1-9, 12, 24, 36, 60] and the values need to be converted before using them.
	if period, ok := d.Get("period").(int); ok && period != 0 {
		request["UsedTime"] = strconv.Itoa(period)
		request["Period"] = Month
		if period > 9 {
			request["UsedTime"] = strconv.Itoa(period / 12)
			request["Period"] = Year
		}
	}
	return request, nil
}
