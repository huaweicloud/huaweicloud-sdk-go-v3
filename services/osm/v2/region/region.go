package region

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"fmt"
)

var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://osm.cn-south-1.myhuaweicloud.cn")

var staticFields = map[string]*region.Region{
	"cn-south-1": CN_SOUTH_1,
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
