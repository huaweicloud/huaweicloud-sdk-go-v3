package region

import (
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://osm.cn-south-1.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-south-1": CN_SOUTH_1,
}

var provider = region.DefaultProviderChain("OSM")

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
