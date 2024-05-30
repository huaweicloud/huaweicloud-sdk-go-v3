package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerNetworkInterfaceResponse Response Object
type ChangeServerNetworkInterfaceResponse struct {

	// 网卡ID。
	Id *string `json:"id,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 网卡IPv4地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 网卡IPv6地址，未开通IPv6协议的网卡不返回该字段。
	Ipv6Address    *string `json:"ipv6_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerNetworkInterfaceResponse", string(data)}, " ")
}
