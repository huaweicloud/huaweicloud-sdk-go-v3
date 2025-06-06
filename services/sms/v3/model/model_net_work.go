package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetWork 网卡实体类
type NetWork struct {

	// 网卡的名称
	Name string `json:"name"`

	// 该网卡绑定的IP
	Ip string `json:"ip"`

	// IPv6地址
	Ipv6 *string `json:"ipv6,omitempty"`

	// 掩码
	Netmask string `json:"netmask"`

	// 网关
	Gateway string `json:"gateway"`

	// Linux必选，网卡的MTU
	Mtu *int32 `json:"mtu,omitempty"`

	// Mac地址
	Mac string `json:"mac"`

	// 数据库ID
	Id *string `json:"id,omitempty"`
}

func (o NetWork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetWork struct{}"
	}

	return strings.Join([]string{"NetWork", string(data)}, " ")
}
