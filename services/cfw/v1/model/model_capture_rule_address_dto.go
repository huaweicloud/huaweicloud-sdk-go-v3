package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureRuleAddressDto struct {

	// 地址
	Address string `json:"address"`

	// 目的地址类型0 ipv4,1 ipv6
	AddressType int32 `json:"address_type"`

	// 输入地址类型，目前只支持0，手工输入类型
	Type int32 `json:"type"`
}

func (o CaptureRuleAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureRuleAddressDto struct{}"
	}

	return strings.Join([]string{"CaptureRuleAddressDto", string(data)}, " ")
}
