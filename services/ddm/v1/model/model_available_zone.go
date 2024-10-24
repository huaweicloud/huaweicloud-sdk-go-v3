package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableZone struct {

	// 可用区CODE
	Code string `json:"code"`

	// 可用区描述
	Description string `json:"description"`

	// 可用区状态
	Status string `json:"status"`

	// 是否支持IPV6
	SupportIpv6 bool `json:"support_ipv6"`
}

func (o AvailableZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}
