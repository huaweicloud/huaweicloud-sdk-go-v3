package region

import (
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://moderation.cn-north-4.myhuaweicloud.com",
		"https://moderation.cn-north-4.myhuaweicloud.cn")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://moderation.cn-north-1.myhuaweicloud.com",
		"https://moderation.cn-north-1.myhuaweicloud.cn")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://moderation.cn-east-3.myhuaweicloud.com",
		"https://moderation.cn-east-3.myhuaweicloud.cn")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://moderation.ap-southeast-1.myhuaweicloud.com",
		"https://moderation.ap-southeast-1.myhuaweicloud.cn")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://moderation.ap-southeast-3.myhuaweicloud.com",
		"https://moderation.ap-southeast-3.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-east-3":      CN_EAST_3,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
}

var provider = region.DefaultProviderChain("MODERATION")

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
