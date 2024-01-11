package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpDetails struct {

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// IPV6地址
	Ipv6Address *string `json:"ipv6_address,omitempty"`

	// 带宽大小
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
}

func (o IpDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpDetails struct{}"
	}

	return strings.Join([]string{"IpDetails", string(data)}, " ")
}
