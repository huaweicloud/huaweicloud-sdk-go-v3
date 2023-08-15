package region

import (
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://devstar.cn-north-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
}

var provider = region.DefaultProviderChain("DEVSTAR")

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
