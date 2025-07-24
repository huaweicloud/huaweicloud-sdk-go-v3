package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://clouddc.cn-east-3.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://clouddc.cn-southwest-2.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://clouddc.cn-south-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://clouddc.ap-southeast-3.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-east-3":      CN_EAST_3,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
}

var provider = region.DefaultProviderChain("CLOUDDC")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'CloudDC': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
