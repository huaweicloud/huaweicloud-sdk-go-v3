package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyProtocolExtension struct {

	// ipv4 vip地址
	VipAddress *string `json:"vip_address,omitempty"`

	// ipv6 vip地址
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	Extension *Extension `json:"extension"`
}

func (o ProxyProtocolExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyProtocolExtension struct{}"
	}

	return strings.Join([]string{"ProxyProtocolExtension", string(data)}, " ")
}
