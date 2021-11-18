package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudPolarDBStoragePlansDataSource(t *testing.T) {
	rand := acctest.RandInt()

	resourceId := "data.alicloud_polardb_storage_plans.default"

	var existPolarDBStoragePlansMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                                      CHECKSET,
			"storage_plans.#":                            CHECKSET,
			"storage_plans.0.id":                         CHECKSET,
			"storage_plans.0.prod_code":                  "polardb",
			"storage_plans.0.ali_uid":                    CHECKSET,
			"storage_plans.0.commodity_code":             "polardb_package",
			"storage_plans.0.template_name":              CHECKSET,
			"storage_plans.0.storage_type":               CHECKSET,
			"storage_plans.0.status":                     CHECKSET,
			"storage_plans.0.start_times":                CHECKSET,
			"storage_plans.0.end_times":                  CHECKSET,
			"storage_plans.0.purchase_times":             CHECKSET,
			"storage_plans.0.init_capacity_view_value":   CHECKSET,
			"storage_plans.0.init_capa_city_view_unit":   CHECKSET,
			"storage_plans.0.period_capacity_view_value": CHECKSET,
			"storage_plans.0.period_capa_city_view_unit": CHECKSET,
			"storage_plans.0.period_time":                CHECKSET,
		}
	}

	var fakePolarDBStoragePlansMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":           "0",
			"storage_plans.#": "0",
		}
	}

	var polarDBStoragePlansCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existPolarDBStoragePlansMapFunc,
		fakeMapFunc:  fakePolarDBStoragePlansMapFunc,
	}

	polarDBStoragePlansCheckInfo.dataSourceTestCheck(t, rand)
}
