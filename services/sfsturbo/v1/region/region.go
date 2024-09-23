package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://sfs-turbo.eu-west-101.myhuaweicloud.eu")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://sfs-turbo.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://sfs-turbo.cn-east-3.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://sfs-turbo.cn-north-1.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://sfs-turbo.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://sfs-turbo.cn-north-4.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://sfs-turbo.cn-north-9.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://sfs-turbo.cn-south-1.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://sfs-turbo.cn-south-2.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://sfs-turbo.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://sfs-turbo.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://sfs-turbo.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://sfs-turbo.ap-southeast-3.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://sfs-turbo.sa-brazil-1.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://sfs-turbo.la-south-2.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://sfs-turbo.la-north-2.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://sfs-turbo.af-south-1.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://sfs-turbo.na-mexico-1.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://sfs-turbo.tr-west-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://sfs-turbo.ap-southeast-4.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://sfs-turbo.me-east-1.myhuaweicloud.com")
	CN_EAST_5 = region.NewRegion("cn-east-5",
		"https://sfs-turbo.cn-east-5.myhuaweicloud.com")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://sfs-turbo.cn-south-4.myhuaweicloud.com")
	RU_MOSCOW_1 = region.NewRegion("ru-moscow-1",
		"https://sfs-turbo.ru-moscow-1.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://sfs-turbo.my-kualalumpur-1.myhuaweicloud.com")
	CN_EAST_4 = region.NewRegion("cn-east-4",
		"https://sfs-turbo.cn-east-4.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://sfs-turbo.ae-ad-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"eu-west-101":      EU_WEST_101,
	"cn-east-2":        CN_EAST_2,
	"cn-east-3":        CN_EAST_3,
	"cn-north-1":       CN_NORTH_1,
	"cn-north-2":       CN_NORTH_2,
	"cn-north-4":       CN_NORTH_4,
	"cn-north-9":       CN_NORTH_9,
	"cn-south-1":       CN_SOUTH_1,
	"cn-south-2":       CN_SOUTH_2,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"sa-brazil-1":      SA_BRAZIL_1,
	"la-south-2":       LA_SOUTH_2,
	"la-north-2":       LA_NORTH_2,
	"af-south-1":       AF_SOUTH_1,
	"na-mexico-1":      NA_MEXICO_1,
	"tr-west-1":        TR_WEST_1,
	"ap-southeast-4":   AP_SOUTHEAST_4,
	"me-east-1":        ME_EAST_1,
	"cn-east-5":        CN_EAST_5,
	"cn-south-4":       CN_SOUTH_4,
	"ru-moscow-1":      RU_MOSCOW_1,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
	"cn-east-4":        CN_EAST_4,
	"ae-ad-1":          AE_AD_1,
}

var provider = region.DefaultProviderChain("SFSTURBO")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'SFSTurbo': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
