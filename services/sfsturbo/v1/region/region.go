package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://sfs-turbo.cn-north-1.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://sfs-turbo.cn-south-2.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://sfs-turbo.cn-north-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://sfs-turbo.cn-east-3.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://sfs-turbo.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://sfs-turbo.cn-south-1.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://sfs-turbo.cn-southwest-2.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://sfs-turbo.cn-north-9.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-1":     CN_NORTH_1,
	"cn-south-2":     CN_SOUTH_2,
	"cn-north-2":     CN_NORTH_2,
	"cn-east-3":      CN_EAST_3,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-north-9":     CN_NORTH_9,
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
