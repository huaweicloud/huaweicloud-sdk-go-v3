package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchIpGroupRequestBodyIpGroups struct {

	// IP地址或网段。支持IPv4、IPv6。
	Ip string `json:"ip"`

	// 备注信息，最长255字符。
	Description *string `json:"description,omitempty"`
}

func (o SwitchIpGroupRequestBodyIpGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupRequestBodyIpGroups struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupRequestBodyIpGroups", string(data)}, " ")
}
