package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://tms.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://tms.ap-southeast-1.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://tms.eu-west-101.myhuaweicloud.eu")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"eu-west-101":    EU_WEST_101,
}

var provider = region.DefaultProviderChain("TMS")

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
