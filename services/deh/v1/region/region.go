package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://deh.cn-north-1.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://deh.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://deh.cn-north-4.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://deh.cn-south-1.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://deh.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://deh.cn-east-3.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://deh.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://deh.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://deh.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://deh.ap-southeast-3.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://deh.my-kualalumpur-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-1":       CN_NORTH_1,
	"cn-north-2":       CN_NORTH_2,
	"cn-north-4":       CN_NORTH_4,
	"cn-south-1":       CN_SOUTH_1,
	"cn-east-2":        CN_EAST_2,
	"cn-east-3":        CN_EAST_3,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
}

var provider = region.DefaultProviderChain("DEH")

func getRegionIds() []string {
	ids := make([]string, 0, len(staticFields))
	for key := range staticFields {
		ids = append(ids, key)
	}
	sort.Strings(ids)
	return ids
}

func SafeValueOf(regionId string) (region *region.Region, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	region = ValueOf(regionId)
	return region, err
}

// Deprecated: This function may panic under certain circumstances. Use SafeValueOf instead.
func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'DeH': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
