package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_EAST_2 = region.NewRegion("cn-east-2", "https://dws.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://dws.cn-east-3.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://dws.cn-north-1.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://dws.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://dws.cn-south-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://dws.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://dws.ap-southeast-3.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-north-1":     CN_NORTH_1,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
}

var provider = region.DefaultProviderChain("DWS")

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
