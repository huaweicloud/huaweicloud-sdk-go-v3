package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://bcs.cn-north-1.myhuaweicloud.com",
		"https://bcs.cn-north-1.myhuaweicloud.cn")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://bcs.cn-north-4.myhuaweicloud.com",
		"https://bcs.cn-north-4.myhuaweicloud.cn")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://bcs.cn-south-1.myhuaweicloud.com",
		"https://bcs.cn-south-1.myhuaweicloud.cn")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://bcs.cn-east-2.myhuaweicloud.com",
		"https://bcs.cn-east-2.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
	"cn-north-4": CN_NORTH_4,
	"cn-south-1": CN_SOUTH_1,
	"cn-east-2":  CN_EAST_2,
}

var provider = region.DefaultProviderChain("BCS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'BCS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
