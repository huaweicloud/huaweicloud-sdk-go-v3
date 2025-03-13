package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://cae.cn-north-4.myhuaweicloud.com",
		"https://cae.cn-north-4.myhuaweicloud.cn")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://cae.cn-east-3.myhuaweicloud.com",
		"https://cae.cn-east-3.myhuaweicloud.cn")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://cae.cn-south-1.myhuaweicloud.com",
		"https://cae.cn-south-1.myhuaweicloud.cn")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://cae.ap-southeast-3.myhuaweicloud.com",
		"https://cae.ap-southeast-3.myhuaweicloud.cn")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://cae.af-south-1.myhuaweicloud.com",
		"https://cae.af-south-1.myhuaweicloud.cn")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://cae.me-east-1.myhuaweicloud.com",
		"https://cae.me-east-1.myhuaweicloud.cn")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://cae.la-north-2.myhuaweicloud.com",
		"https://cae.la-north-2.myhuaweicloud.cn")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://cae.tr-west-1.myhuaweicloud.com",
		"https://cae.tr-west-1.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"af-south-1":     AF_SOUTH_1,
	"me-east-1":      ME_EAST_1,
	"la-north-2":     LA_NORTH_2,
	"tr-west-1":      TR_WEST_1,
}

var provider = region.DefaultProviderChain("CAE")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'CAE': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
