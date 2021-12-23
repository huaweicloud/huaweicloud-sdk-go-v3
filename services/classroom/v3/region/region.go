package region

import (
	"fmt"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://classroom-ext.cn-north-4.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4": CN_NORTH_4,
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
