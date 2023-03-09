package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://bcs.cn-north-1.myhuaweicloud.com",
		"https://bcs.cn-north-1.myhuaweicloud.cn")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://bcs.cn-north-4.myhuaweicloud.com",
		"https://bcs.cn-north-4.myhuaweicloud.cn")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://bcs.cn-south-1.myhuaweicloud.com",
		"https://bcs.cn-south-1.myhuaweicloud.cn")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://bcs.cn-east-2.myhuaweicloud.com",
		"https://bcs.cn-east-2.myhuaweicloud.cn")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://bcs.ap-southeast-1.myhuaweicloud.com",
		"https://bcs.ap-southeast-1.myhuaweicloud.cn")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://bcs.ap-southeast-2.myhuaweicloud.com",
		"https://bcs.ap-southeast-2.myhuaweicloud.cn")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://bcs.ap-southeast-3.myhuaweicloud.com",
		"https://bcs.ap-southeast-3.myhuaweicloud.cn")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://bcs.ap-southeast-4.myhuaweicloud.com",
		"https://bcs.ap-southeast-4.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-north-1":     CN_NORTH_1,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-2":      CN_EAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-4": AP_SOUTHEAST_4,
}

var provider = region.DefaultProviderChain("BCS")

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
