package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudPolarDBStoragePlanUpdate(t *testing.T) {
	var v map[string]interface{}
	name := "tf-testAccDBStoragePlanUpdate"
	resourceId := "alicloud_polardb_storage_plan.default"
	ra := resourceAttrInit(resourceId, storagePlanBasicMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDBService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDBStoragePlan")
	rac := resourceAttrCheckInit(rc, ra)

	testAccCheck := rac.resourceAttrMapUpdateSet()
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceDBStoragePlanConfigDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		//因为没有API delete
		//CheckDestroy: rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_type":  "Overseas",
					"period":        "1",
					"storage_class": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_type":  CHECKSET,
						"storage_class": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period", "storage_class"},
			},
		},
	})

}

var storagePlanBasicMap = map[string]string{
	"storage_type":   CHECKSET,
	"period":         CHECKSET,
	"storage_class":  CHECKSET,
	"prod_code":      "polardb",
	"commodity_code": "polardb_package",
}

func resourceDBStoragePlanConfigDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "%s"
	}
	`, name)
}
