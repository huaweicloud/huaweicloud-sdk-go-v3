package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://lakeformation.cn-north-4.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://lakeformation.cn-north-9.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://lakeformation.cn-east-3.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://lakeformation.cn-south-1.myhuaweicloud.com")
	CN_NORTH_11 = region.NewRegion("cn-north-11",
		"https://lakeformation.cn-north-11.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://lakeformation.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://lakeformation.ap-southeast-4.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://lakeformation.la-north-2.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://lakeformation.af-south-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-9":     CN_NORTH_9,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"cn-north-11":    CN_NORTH_11,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-4": AP_SOUTHEAST_4,
	"la-north-2":     LA_NORTH_2,
	"af-south-1":     AF_SOUTH_1,
}

var provider = region.DefaultProviderChain("LAKEFORMATION")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'LakeFormation': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
