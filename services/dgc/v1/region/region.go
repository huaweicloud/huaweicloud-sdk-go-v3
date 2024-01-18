package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://dayu-dlf.cn-north-1.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://dayu-dlf.cn-north-4.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://dayu-dlf.cn-east-3.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://dayu-dlf.cn-east-2.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://dayu-dlf.cn-south-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://dayu-dlf.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://dayu-dlf.ap-southeast-3.myhuaweicloud.com")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://dayu-dlf.ru-northwest-2.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-1":     CN_NORTH_1,
	"cn-north-4":     CN_NORTH_4,
	"cn-east-3":      CN_EAST_3,
	"cn-east-2":      CN_EAST_2,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ru-northwest-2": RU_NORTHWEST_2,
}

var provider = region.DefaultProviderChain("DGC")

func getRegionIds() []string {
	ids := make([]string, 0, len(staticFields))
	for key := range staticFields {
		ids = append(ids, key)
	}
	sort.Strings(ids)
	return ids
}

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'DGC': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
