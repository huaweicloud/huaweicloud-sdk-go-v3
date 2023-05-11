package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NovaServerInterfaceDetail struct {

	// 网卡私网IP信息列表。
	FixedIps []NovaServerInterfaceFixedIp `json:"fixed_ips"`

	// 网卡Mac地址信息。
	MacAddr string `json:"mac_addr"`

	// 网卡端口所属网络ID。
	NetId string `json:"net_id"`

	// 网卡端口ID。
	PortId string `json:"port_id"`

	// 网卡端口状态。
	PortState string `json:"port_state"`
}

func (o NovaServerInterfaceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerInterfaceDetail struct{}"
	}

	return strings.Join([]string{"NovaServerInterfaceDetail", string(data)}, " ")
}
