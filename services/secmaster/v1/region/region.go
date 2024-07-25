package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://secmaster.cn-north-4.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://secmaster.cn-east-3.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://secmaster.cn-south-1.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://secmaster.cn-north-9.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://secmaster.cn-north-2.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://secmaster.cn-east-2.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://secmaster.cn-southwest-2.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://secmaster.cn-south-2.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://secmaster.cn-north-1.myhuaweicloud.com")
	CN_EAST_4 = region.NewRegion("cn-east-4",
		"https://secmaster.cn-east-4.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://secmaster.me-east-1.myhuaweicloud.com")
	CN_NORTH_11 = region.NewRegion("cn-north-11",
		"https://secmaster.cn-north-11.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://secmaster.sa-brazil-1.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://secmaster.la-north-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://secmaster.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://secmaster.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://secmaster.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://secmaster.ap-southeast-4.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://secmaster.tr-west-1.myhuaweicloud.com")
	AF_NORTH_1 = region.NewRegion("af-north-1",
		"https://secmaster.af-north-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"cn-north-9":     CN_NORTH_9,
	"cn-north-2":     CN_NORTH_2,
	"cn-east-2":      CN_EAST_2,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-2":     CN_SOUTH_2,
	"cn-north-1":     CN_NORTH_1,
	"cn-east-4":      CN_EAST_4,
	"me-east-1":      ME_EAST_1,
	"cn-north-11":    CN_NORTH_11,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-4": AP_SOUTHEAST_4,
	"tr-west-1":      TR_WEST_1,
	"af-north-1":     AF_NORTH_1,
}

var provider = region.DefaultProviderChain("SECMASTER")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'SecMaster': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
