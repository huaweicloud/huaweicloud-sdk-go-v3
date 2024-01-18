package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://cbs-ext.cn-north-1.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://cbs-ext.cn-north-4.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-1": CN_NORTH_1,
	"cn-north-4": CN_NORTH_4,
}

var provider = region.DefaultProviderChain("CBS")

func getRegionIds() []string {
	ids := make([]string, 0, len(staticFields))
	for key := range staticFields {
		ids = append(ids, key)
	}
	sort.Strings(ids)
	return ids
}

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'CBS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
