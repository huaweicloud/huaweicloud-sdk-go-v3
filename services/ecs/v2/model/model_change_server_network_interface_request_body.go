package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerNetworkInterfaceRequestBody This is a auto create Body Object
type ChangeServerNetworkInterfaceRequestBody struct {

	// 子网ID ，UUID格式。更新IP地址时，必须指定该参数；更新IPv6地址时，该参数可以不填。
	SubnetId *string `json:"subnet_id,omitempty"`

	// IPv4地址，为空字符串时表示随机更新网卡IPv4地址。  约束： - 该参数和ipv6_address必须指定一个。 - 该参数和ipv6_address不能同时指定。
	IpAddress *string `json:"ip_address,omitempty"`

	// IPv6地址，为空字符串时表示随机更新网卡IPv6地址。  约束： - 该参数和ip_address必须指定一个。 - 该参数和ip_address不能同时指定。
	Ipv6Address *string `json:"ipv6_address,omitempty"`
}

func (o ChangeServerNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeServerNetworkInterfaceRequestBody", string(data)}, " ")
}
