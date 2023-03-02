package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://codecheck-ext.cn-north-4.myhuaweicloud.com",
		"https://codecheck-ext.cn-north-4.myhuaweicloud.cn")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://codecheck-ext.cn-north-1.myhuaweicloud.com",
		"https://codecheck-ext.cn-north-1.myhuaweicloud.cn")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://codecheck-ext.cn-east-2.myhuaweicloud.com",
		"https://codecheck-ext.cn-east-2.myhuaweicloud.cn")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://codecheck-ext.cn-south-1.myhuaweicloud.com",
		"https://codecheck-ext.cn-south-1.myhuaweicloud.cn")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://codecheck-ext.cn-east-3.myhuaweicloud.com",
		"https://codecheck-ext.cn-east-3.myhuaweicloud.cn")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://codecheck-ext.la-north-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://codecheck-ext.sa-brazil-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://codecheck-ext.ap-southeast-3.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-east-2":      CN_EAST_2,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"la-north-2":     LA_NORTH_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
}

var provider = region.DefaultProviderChain("CODECHECK")

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
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
