package region

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/region"
	"fmt"
)

var CN_EAST_2 = region.NewRegion("cn-east-2", "https://vod.cn-east-2.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://vod.cn-north-1.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://vod.cn-north-4.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-east-2":  CN_EAST_2,
	"cn-north-1": CN_NORTH_1,
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
