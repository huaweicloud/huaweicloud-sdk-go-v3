package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://ocr.cn-north-4.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://ocr.cn-south-1.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://ocr.cn-east-3.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://ocr.cn-north-1.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://ocr.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://ocr.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://ocr.ap-southeast-3.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://ocr.la-south-2.myhuaweicloud.com")
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://ocr.af-south-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-north-1":     CN_NORTH_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"la-south-2":     LA_SOUTH_2,
	"af-south-1":     AF_SOUTH_1,
}

var provider = region.DefaultProviderChain("OCR")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'OCR': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
