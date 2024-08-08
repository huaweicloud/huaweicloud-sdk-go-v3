package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://iotda.cn-north-4.myhuaweicloud.com")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://iotda.cn-south-4.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://iotda.cn-south-1.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://iotda.cn-east-3.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://iotda.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://iotda.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://iotda.ap-southeast-1.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://iotda.af-south-1.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://iotda.me-east-1.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://iotda.la-south-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://iotda.sa-brazil-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-south-4":     CN_SOUTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"af-south-1":     AF_SOUTH_1,
	"me-east-1":      ME_EAST_1,
	"la-south-2":     LA_SOUTH_2,
	"sa-brazil-1":    SA_BRAZIL_1,
}

var provider = region.DefaultProviderChain("IOTDA")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'IoTDA': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
