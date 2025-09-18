package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudConnectionIpv6SupportRegionsCapabilityValue struct {

	// 租户支持的区域列表。
	SupportRegions []string `json:"support_regions"`
}

func (o CloudConnectionIpv6SupportRegionsCapabilityValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionIpv6SupportRegionsCapabilityValue struct{}"
	}

	return strings.Join([]string{"CloudConnectionIpv6SupportRegionsCapabilityValue", string(data)}, " ")
}
