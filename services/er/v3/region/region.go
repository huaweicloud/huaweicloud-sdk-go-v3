package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://er.cn-south-1.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://er.cn-north-9.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://er.ap-southeast-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-1":     CN_SOUTH_1,
	"cn-north-9":     CN_NORTH_9,
	"ap-southeast-2": AP_SOUTHEAST_2,
}

var provider = region.DefaultProviderChain("ER")

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
