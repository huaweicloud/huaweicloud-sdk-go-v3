package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://dme.cn-north-4.myhuaweicloud.com")
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://dme.cn-south-4.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4": CN_NORTH_4,
	"cn-south-4": CN_SOUTH_4,
}

var provider = region.DefaultProviderChain("IDME")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'iDME': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
