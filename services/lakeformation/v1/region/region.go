package region

import (
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://lakeformation.cn-north-9.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://lakeformation.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://lakeformation.cn-north-4.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-9": CN_NORTH_9,
	"cn-north-2": CN_NORTH_2,
	"cn-north-4": CN_NORTH_4,
}

var provider = region.DefaultProviderChain("LAKEFORMATION")

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
