package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_4 = region.NewRegion("cn-south-4", "https://cfw.cn-south-4.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://cfw.cn-southwest-2.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://cfw.cn-south-1.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://cfw.cn-east-3.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://cfw.cn-north-4.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-4":     CN_SOUTH_4,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-north-4":     CN_NORTH_4,
}

var provider = region.DefaultProviderChain("CFW")

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
