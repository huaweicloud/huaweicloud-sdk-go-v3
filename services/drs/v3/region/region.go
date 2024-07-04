package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://drs.eu-west-101.myhuaweicloud.eu")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://drs.cn-north-4.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://drs.cn-north-1.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://drs.cn-south-1.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://drs.cn-east-3.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://drs.cn-east-2.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://drs.cn-north-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://drs.ap-southeast-3.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://drs.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://drs.ap-southeast-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://drs.sa-brazil-1.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://drs.la-south-2.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://drs.la-north-2.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://drs.na-mexico-1.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://drs.af-south-1.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://drs.cn-north-9.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://drs.cn-southwest-2.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://drs.cn-south-2.myhuaweicloud.com")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://drs.cn-south-4.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://drs.tr-west-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://drs.ap-southeast-4.myhuaweicloud.com")
	EU_WEST_0 = region.NewRegion("eu-west-0",
		"https://drs.eu-west-0.myhuaweicloud.com")
	RU_MOSCOW_1 = region.NewRegion("ru-moscow-1",
		"https://drs.ru-moscow-1.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://drs.ae-ad-1.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://drs.my-kualalumpur-1.myhuaweicloud.com")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://drs.ru-northwest-2.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"eu-west-101":      EU_WEST_101,
	"cn-north-4":       CN_NORTH_4,
	"cn-north-1":       CN_NORTH_1,
	"cn-south-1":       CN_SOUTH_1,
	"cn-east-3":        CN_EAST_3,
	"cn-east-2":        CN_EAST_2,
	"cn-north-2":       CN_NORTH_2,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"sa-brazil-1":      SA_BRAZIL_1,
	"la-south-2":       LA_SOUTH_2,
	"la-north-2":       LA_NORTH_2,
	"na-mexico-1":      NA_MEXICO_1,
	"af-south-1":       AF_SOUTH_1,
	"cn-north-9":       CN_NORTH_9,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"cn-south-2":       CN_SOUTH_2,
	"cn-south-4":       CN_SOUTH_4,
	"tr-west-1":        TR_WEST_1,
	"ap-southeast-4":   AP_SOUTHEAST_4,
	"eu-west-0":        EU_WEST_0,
	"ru-moscow-1":      RU_MOSCOW_1,
	"ae-ad-1":          AE_AD_1,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
	"ru-northwest-2":   RU_NORTHWEST_2,
}

var provider = region.DefaultProviderChain("DRS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'DRS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
