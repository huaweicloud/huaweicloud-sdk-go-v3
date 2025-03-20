package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowAddressNetmask 允许访问的IP地址或网段。
type AllowAddressNetmask struct {

	// IP地址或网段，例如\"192.168.0.1/24\"。
	AddressNetmask string `json:"address_netmask"`

	// 描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`
}

func (o AllowAddressNetmask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowAddressNetmask struct{}"
	}

	return strings.Join([]string{"AllowAddressNetmask", string(data)}, " ")
}
