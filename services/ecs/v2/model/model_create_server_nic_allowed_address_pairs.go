package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServerNicAllowedAddressPairs struct {

	// IP地址 不支持0.0.0.0/0 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。
	IpAddress *string `json:"ip_address,omitempty"`

	// MAC地址
	MacAddress *string `json:"mac_address,omitempty"`
}

func (o CreateServerNicAllowedAddressPairs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerNicAllowedAddressPairs struct{}"
	}

	return strings.Join([]string{"CreateServerNicAllowedAddressPairs", string(data)}, " ")
}
