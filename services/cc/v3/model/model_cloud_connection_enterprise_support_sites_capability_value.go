package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudConnectionEnterpriseSupportSitesCapabilityValue struct {

	// 租户支持的Site列表。
	SupportSites []string `json:"support_sites"`
}

func (o CloudConnectionEnterpriseSupportSitesCapabilityValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionEnterpriseSupportSitesCapabilityValue struct{}"
	}

	return strings.Join([]string{"CloudConnectionEnterpriseSupportSitesCapabilityValue", string(data)}, " ")
}
