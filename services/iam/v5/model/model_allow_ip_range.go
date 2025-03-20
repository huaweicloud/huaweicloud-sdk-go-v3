package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowIpRange 允许访问的IP地址区间。
type AllowIpRange struct {

	// IP地址区间，例如\"0.0.0.0-255.255.255.255\"。
	IpRange string `json:"ip_range"`

	// 描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`
}

func (o AllowIpRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowIpRange struct{}"
	}

	return strings.Join([]string{"AllowIpRange", string(data)}, " ")
}
