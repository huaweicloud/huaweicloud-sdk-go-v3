package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://rfs.cn-north-4.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://rfs.cn-south-1.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://rfs.cn-east-3.myhuaweicloud.com")
	CN_EAST_4 = region.NewRegion("cn-east-4",
		"https://rfs.cn-east-4.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://rfs.cn-north-9.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://rfs.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://rfs.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://rfs.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://rfs.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://rfs.ap-southeast-4.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://rfs.me-east-1.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://rfs.tr-west-1.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://rfs.la-south-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://rfs.sa-brazil-1.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://rfs.la-north-2.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://rfs.af-south-1.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://aos.myhuaweicloud.eu")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://rfs.na-mexico-1.myhuaweicloud.com")
	CN_NORTH_11 = region.NewRegion("cn-north-11",
		"https://rfs.cn-north-11.myhuaweicloud.com")
	CN_EAST_5 = region.NewRegion("cn-east-5",
		"https://rfs.cn-east-5.myhuaweicloud.com")
	AF_NORTH_1 = region.NewRegion("af-north-1",
		"https://rfs.af-north-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-east-4":      CN_EAST_4,
	"cn-north-9":     CN_NORTH_9,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-4": AP_SOUTHEAST_4,
	"me-east-1":      ME_EAST_1,
	"tr-west-1":      TR_WEST_1,
	"la-south-2":     LA_SOUTH_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
	"af-south-1":     AF_SOUTH_1,
	"eu-west-101":    EU_WEST_101,
	"na-mexico-1":    NA_MEXICO_1,
	"cn-north-11":    CN_NORTH_11,
	"cn-east-5":      CN_EAST_5,
	"af-north-1":     AF_NORTH_1,
}

var provider = region.DefaultProviderChain("AOS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'AOS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
