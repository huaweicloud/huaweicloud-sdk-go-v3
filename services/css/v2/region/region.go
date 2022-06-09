package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://css.cn-south-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-1": CN_SOUTH_1,
}

var provider = region.DefaultProviderChain("CSS")

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
