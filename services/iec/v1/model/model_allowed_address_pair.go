package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowedAddressPair IP/Mac对象
type AllowedAddressPair struct {

	// - 功能说明：IP地址 - 约束：     IP地址不支持“0.0.0.0/0”     如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。     如果allowed_address_pairs的IP地址为“1.1.1.1/0”，表示关闭源目地址检查开关。
	IpAddress string `json:"ip_address"`

	// MAC地址
	MacAddress *string `json:"mac_address,omitempty"`
}

func (o AllowedAddressPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowedAddressPair struct{}"
	}

	return strings.Join([]string{"AllowedAddressPair", string(data)}, " ")
}
