package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowIpRangesResultIPv6
type AllowIpRangesResultIPv6 struct {

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// IPv6地址区间，例如： 0000:0000:0000:0000:0000:0000:0000:0000-FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF。
	IpRange string `json:"ip_range"`
}

func (o AllowIpRangesResultIPv6) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowIpRangesResultIPv6 struct{}"
	}

	return strings.Join([]string{"AllowIpRangesResultIPv6", string(data)}, " ")
}
