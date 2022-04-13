package region

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/region"
	"fmt"
)

var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://ccaas.cn-north-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
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
