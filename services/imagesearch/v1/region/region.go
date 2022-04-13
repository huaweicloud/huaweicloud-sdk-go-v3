package region

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"fmt"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://imagesearch.cn-north-4.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://imagesearch.cn-north-1.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://imagesearch.ap-southeast-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
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
