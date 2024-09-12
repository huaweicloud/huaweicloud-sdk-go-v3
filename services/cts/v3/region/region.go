package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://cts.af-south-1.myhuaweicloud.com",
		"https://cts.af-south-1.myhuaweicloud.cn")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://cts.cn-north-4.myhuaweicloud.com",
		"https://cts.cn-north-4.myhuaweicloud.cn")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://cts.cn-north-1.myhuaweicloud.com",
		"https://cts.cn-north-1.myhuaweicloud.cn")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://cts.cn-east-2.myhuaweicloud.com",
		"https://cts.cn-east-2.myhuaweicloud.cn")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://cts.cn-east-3.myhuaweicloud.com",
		"https://cts.cn-east-3.myhuaweicloud.cn")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://cts.cn-south-1.myhuaweicloud.com",
		"https://cts.cn-south-1.myhuaweicloud.cn")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://cts.cn-south-2.myhuaweicloud.com",
		"https://cts.cn-south-2.myhuaweicloud.cn")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://cts.cn-southwest-2.myhuaweicloud.com",
		"https://cts.cn-southwest-2.myhuaweicloud.cn")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://cts.ap-southeast-2.myhuaweicloud.com",
		"https://cts.ap-southeast-2.myhuaweicloud.cn")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://cts.ap-southeast-1.myhuaweicloud.com",
		"https://cts.ap-southeast-1.myhuaweicloud.cn")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://cts.ap-southeast-3.myhuaweicloud.com",
		"https://cts.ap-southeast-3.myhuaweicloud.cn")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://cts.cn-north-2.myhuaweicloud.com",
		"https://cts.cn-north-2.myhuaweicloud.cn")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://cts.cn-north-9.myhuaweicloud.com",
		"https://cts.cn-north-9.myhuaweicloud.cn")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://cts.cn-south-4.myhuaweicloud.com",
		"https://cts.cn-south-4.myhuaweicloud.cn")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://cts.ru-northwest-2.myhuaweicloud.com",
		"https://cts.ru-northwest-2.myhuaweicloud.cn")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://cts.la-south-2.myhuaweicloud.com",
		"https://cts.la-south-2.myhuaweicloud.cn")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://cts.sa-brazil-1.myhuaweicloud.com",
		"https://cts.sa-brazil-1.myhuaweicloud.cn")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://cts.la-north-2.myhuaweicloud.com",
		"https://cts.la-north-2.myhuaweicloud.cn")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://cts.na-mexico-1.myhuaweicloud.com",
		"https://cts.na-mexico-1.myhuaweicloud.cn")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://cts.eu-west-101.myhuaweicloud.eu")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://cts.ap-southeast-4.myhuaweicloud.com",
		"https://cts.ap-southeast-4.myhuaweicloud.cn")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://cts.tr-west-1.myhuaweicloud.com",
		"https://cts.tr-west-1.myhuaweicloud.cn")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://cts.me-east-1.myhuaweicloud.com")
	EU_WEST_0 = region.NewRegion("eu-west-0",
		"https://cts.eu-west-0.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://cts.my-kualalumpur-1.myhuaweicloud.com")
	RU_MOSCOW_1 = region.NewRegion("ru-moscow-1",
		"https://cts.ru-moscow-1.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://cts.ae-ad-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"af-south-1":       AF_SOUTH_1,
	"cn-north-4":       CN_NORTH_4,
	"cn-north-1":       CN_NORTH_1,
	"cn-east-2":        CN_EAST_2,
	"cn-east-3":        CN_EAST_3,
	"cn-south-1":       CN_SOUTH_1,
	"cn-south-2":       CN_SOUTH_2,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"cn-north-2":       CN_NORTH_2,
	"cn-north-9":       CN_NORTH_9,
	"cn-south-4":       CN_SOUTH_4,
	"ru-northwest-2":   RU_NORTHWEST_2,
	"la-south-2":       LA_SOUTH_2,
	"sa-brazil-1":      SA_BRAZIL_1,
	"la-north-2":       LA_NORTH_2,
	"na-mexico-1":      NA_MEXICO_1,
	"eu-west-101":      EU_WEST_101,
	"ap-southeast-4":   AP_SOUTHEAST_4,
	"tr-west-1":        TR_WEST_1,
	"me-east-1":        ME_EAST_1,
	"eu-west-0":        EU_WEST_0,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
	"ru-moscow-1":      RU_MOSCOW_1,
	"ae-ad-1":          AE_AD_1,
}

var provider = region.DefaultProviderChain("CTS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'CTS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
