package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_EAST_3 = region.NewRegion("cn-east-3", "https://cloudartifacts-ext.cn-east-3.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-east-3": CN_EAST_3,
}

var provider = region.DefaultProviderChain("CLOUDARTIFACT")

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
