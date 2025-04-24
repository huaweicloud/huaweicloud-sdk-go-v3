package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://bms.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://bms.cn-north-4.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://bms.cn-north-1.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://bms.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://bms.cn-east-3.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://bms.cn-south-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://bms.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://bms.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://bms.ap-southeast-1.myhuaweicloud.com")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://bms.ru-northwest-2.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://bms.la-south-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://bms.sa-brazil-1.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://bms.cn-north-9.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://bms.cn-southwest-2.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://bms.af-south-1.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://bms.la-north-2.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://bms.cn-south-2.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://bms.na-mexico-1.myhuaweicloud.com")
	CN_EAST_5 = region.NewRegion("cn-east-5",
		"https://bms.cn-east-5.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://bms.ap-southeast-4.myhuaweicloud.com")
	CN_NORTH_11 = region.NewRegion("cn-north-11",
		"https://bms.cn-north-11.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://bms.tr-west-1.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://bms.me-east-1.myhuaweicloud.com")
	CN_EAST_4 = region.NewRegion("cn-east-4",
		"https://bms.cn-east-4.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://bms.my-kualalumpur-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-2":       CN_NORTH_2,
	"cn-north-4":       CN_NORTH_4,
	"cn-north-1":       CN_NORTH_1,
	"cn-east-2":        CN_EAST_2,
	"cn-east-3":        CN_EAST_3,
	"cn-south-1":       CN_SOUTH_1,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ru-northwest-2":   RU_NORTHWEST_2,
	"la-south-2":       LA_SOUTH_2,
	"sa-brazil-1":      SA_BRAZIL_1,
	"cn-north-9":       CN_NORTH_9,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"af-south-1":       AF_SOUTH_1,
	"la-north-2":       LA_NORTH_2,
	"cn-south-2":       CN_SOUTH_2,
	"na-mexico-1":      NA_MEXICO_1,
	"cn-east-5":        CN_EAST_5,
	"ap-southeast-4":   AP_SOUTHEAST_4,
	"cn-north-11":      CN_NORTH_11,
	"tr-west-1":        TR_WEST_1,
	"me-east-1":        ME_EAST_1,
	"cn-east-4":        CN_EAST_4,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
}

var provider = region.DefaultProviderChain("BMS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'BMS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
