package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnGatewayRoutingTableEntryVo struct {

	// 目的地址
	Destination *string `json:"destination,omitempty"`

	// 下一跳地址
	Nexthop *string `json:"nexthop,omitempty"`

	// 出接口地址
	OutboundInterfaceIp *string `json:"outbound_interface_ip,omitempty"`

	// BGP路由来源
	Origin *string `json:"origin,omitempty"`

	// BGP路由的AS路径
	AsPath *string `json:"as_path,omitempty"`

	// BGP路由的MED值
	Med *int64 `json:"med,omitempty"`

	NexthopResource *NexthopResourceVo `json:"nexthop_resource,omitempty"`
}

func (o VpnGatewayRoutingTableEntryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnGatewayRoutingTableEntryVo struct{}"
	}

	return strings.Join([]string{"VpnGatewayRoutingTableEntryVo", string(data)}, " ")
}
