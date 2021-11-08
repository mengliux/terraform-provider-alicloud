package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudPolarDBStoragePlansDataSource(t *testing.T) {
	rand := acctest.RandInt()

	resourceId := "data.alicloud_polardb_storage_plans.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId, "", dataSourcePolarDBStoragePlansConfigDependence)

	prodCodeConfig := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"prod_code": "polardb",
		}),
	}

	commodityCodeConfig := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"commodity_code": "polardb_package",
		}),
	}

	statusConfig := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"status": "valid",
		}),
	}

	var existPolarDBStoragePlansMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                                                     CHECKSET,
			"describeStoragePlanResponses.#":                            CHECKSET,
			"describeStoragePlanResponses.0.id":                         CHECKSET,
			"describeStoragePlanResponses.0.prod_code":                  "polardb",
			"describeStoragePlanResponses.0.ali_uid":                    CHECKSET,
			"describeStoragePlanResponses.0.commodity_code":             "polardb_package",
			"describeStoragePlanResponses.0.template_name":              CHECKSET,
			"describeStoragePlanResponses.0.storage_type":               CHECKSET,
			"describeStoragePlanResponses.0.other_property":             CHECKSET,
			"describeStoragePlanResponses.0.status":                     CHECKSET,
			"describeStoragePlanResponses.0.start_times":                CHECKSET,
			"describeStoragePlanResponses.0.end_times":                  CHECKSET,
			"describeStoragePlanResponses.0.purchase_times":             CHECKSET,
			"describeStoragePlanResponses.0.init_capacity_view_value":   CHECKSET,
			"describeStoragePlanResponses.0.init_capa_city_view_unit":   CHECKSET,
			"describeStoragePlanResponses.0.period_capacity_view_value": CHECKSET,
			"describeStoragePlanResponses.0.period_capa_city_view_unit": CHECKSET,
			"describeStoragePlanResponses.0.period_time":                CHECKSET,
		}
	}

	var fakePolarDBStoragePlansMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                          "0",
			"describeStoragePlanResponses.#": "0",
		}
	}

	var polarDBStoragePlansCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existPolarDBStoragePlansMapFunc,
		fakeMapFunc:  fakePolarDBStoragePlansMapFunc,
	}

	polarDBStoragePlansCheckInfo.dataSourceTestCheck(t, rand, prodCodeConfig, commodityCodeConfig, statusConfig)
}

func dataSourcePolarDBStoragePlansConfigDependence(name string) string {
	return ""
}
