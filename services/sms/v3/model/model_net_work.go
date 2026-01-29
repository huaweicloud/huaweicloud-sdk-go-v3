package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetWork 网卡实体类
type NetWork struct {

	// 网卡的名称
	Name *string `json:"name,omitempty"`

	// 该网卡绑定的IP
	Ip *string `json:"ip,omitempty"`

	// IPv6地址
	Ipv6 *string `json:"ipv6,omitempty"`

	// 掩码
	Netmask *string `json:"netmask,omitempty"`

	// 网关
	Gateway *string `json:"gateway,omitempty"`

	// Linux必选，网卡的MTU
	Mtu *int32 `json:"mtu,omitempty"`

	// 列表中第一个Mac地址必须非空
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
