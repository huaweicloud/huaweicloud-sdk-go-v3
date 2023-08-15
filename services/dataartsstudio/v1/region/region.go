package region

import (
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://dayu.cn-north-1.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://dayu.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://dayu.cn-north-4.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://dayu.cn-east-3.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://dayu.cn-east-2.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://dayu.cn-south-1.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://dayu.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://dayu.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://dayu.ap-southeast-3.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://dayu.af-south-1.myhuaweicloud.com")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://dayu.ru-northwest-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://dayu.sa-brazil-1.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://dayu.la-south-2.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://dayu.la-north-2.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://dayu.na-mexico-1.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://dayu.cn-north-9.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://dayu.cn-southwest-2.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-1":     CN_NORTH_1,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-east-3":      CN_EAST_3,
	"cn-east-2":      CN_EAST_2,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"af-south-1":     AF_SOUTH_1,
	"ru-northwest-2": RU_NORTHWEST_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-south-2":     LA_SOUTH_2,
	"la-north-2":     LA_NORTH_2,
	"na-mexico-1":    NA_MEXICO_1,
	"cn-north-9":     CN_NORTH_9,
	"cn-southwest-2": CN_SOUTHWEST_2,
}

var provider = region.DefaultProviderChain("DATAARTSSTUDIO")

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
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
