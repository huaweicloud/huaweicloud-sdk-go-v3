package region

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"fmt"
)

var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://sis-ext.cn-north-1.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://sis-ext.cn-north-4.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://sis-ext.cn-east-3.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
	"cn-north-4": CN_NORTH_4,
	"cn-east-3":  CN_EAST_3,
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
