package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://ims.af-south-1.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://ims.cn-north-4.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://ims.cn-north-1.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://ims.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://ims.cn-east-3.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://ims.cn-south-1.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://ims.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://ims.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://ims.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://ims.ap-southeast-3.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://ims.cn-north-2.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://ims.cn-south-2.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://ims.cn-north-9.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://ims.la-south-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://ims.sa-brazil-1.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://ims.la-north-2.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://ims.na-mexico-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://ims.ap-southeast-4.myhuaweicloud.com")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://ims.cn-south-4.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://ims.tr-west-1.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://ims.me-east-1.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://ims.ae-ad-1.myhuaweicloud.com")
	CN_EAST_4 = region.NewRegion("cn-east-4",
		"https://ims.cn-east-4.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://ims.eu-west-101.myhuaweicloud.com")
	CN_EAST_5 = region.NewRegion("cn-east-5",
		"https://ims.cn-east-5.myhuaweicloud.com")
	EU_WEST_0 = region.NewRegion("eu-west-0",
		"https://ims.eu-west-0.myhuaweicloud.com")
	MY_KUALALUMPUR_1 = region.NewRegion("my-kualalumpur-1",
		"https://ims.my-kualalumpur-1.myhuaweicloud.com")
	AF_NORTH_1 = region.NewRegion("af-north-1",
		"https://ims.af-north-1.myhuaweicloud.com")
	RU_MOSCOW_1 = region.NewRegion("ru-moscow-1",
		"https://ims.ru-moscow-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"af-south-1":       AF_SOUTH_1,
	"cn-north-4":       CN_NORTH_4,
	"cn-north-1":       CN_NORTH_1,
	"cn-east-2":        CN_EAST_2,
	"cn-east-3":        CN_EAST_3,
	"cn-south-1":       CN_SOUTH_1,
	"cn-southwest-2":   CN_SOUTHWEST_2,
	"ap-southeast-2":   AP_SOUTHEAST_2,
	"ap-southeast-1":   AP_SOUTHEAST_1,
	"ap-southeast-3":   AP_SOUTHEAST_3,
	"cn-north-2":       CN_NORTH_2,
	"cn-south-2":       CN_SOUTH_2,
	"cn-north-9":       CN_NORTH_9,
	"la-south-2":       LA_SOUTH_2,
	"sa-brazil-1":      SA_BRAZIL_1,
	"la-north-2":       LA_NORTH_2,
	"na-mexico-1":      NA_MEXICO_1,
	"ap-southeast-4":   AP_SOUTHEAST_4,
	"cn-south-4":       CN_SOUTH_4,
	"tr-west-1":        TR_WEST_1,
	"me-east-1":        ME_EAST_1,
	"ae-ad-1":          AE_AD_1,
	"cn-east-4":        CN_EAST_4,
	"eu-west-101":      EU_WEST_101,
	"cn-east-5":        CN_EAST_5,
	"eu-west-0":        EU_WEST_0,
	"my-kualalumpur-1": MY_KUALALUMPUR_1,
	"af-north-1":       AF_NORTH_1,
	"ru-moscow-1":      RU_MOSCOW_1,
}

var provider = region.DefaultProviderChain("IMS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'IMS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
