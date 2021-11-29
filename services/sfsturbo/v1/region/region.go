package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://sfs-turbo.cn-north-1.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://sfs-turbo.cn-south-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
	"cn-south-2": CN_SOUTH_2,
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
