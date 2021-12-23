package region

import (
	"fmt"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://face.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://face.cn-south-1.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://face.cn-east-3.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://face.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://face.ap-southeast-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
}

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}
	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
