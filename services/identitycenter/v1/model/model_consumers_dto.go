package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsumersDto SAML响应接收方
type ConsumersDto struct {

	// SAML传输协议
	Binding string `json:"binding"`

	// 是否是默认接收方
	DefaultValue bool `json:"default_value"`

	// SAML ACS Url
	Location string `json:"location"`
}

func (o ConsumersDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumersDto struct{}"
	}

	return strings.Join([]string{"ConsumersDto", string(data)}, " ")
}
