package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ips struct {

	// ip id
	IpId *string `json:"ip_id,omitempty"`

	// ip
	Ip *string `json:"ip,omitempty"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// 数据中心
	DataCenter *string `json:"data_center,omitempty"`

	// 海外区域封禁状态 0-关闭 1-开启
	ForeignSwitchStatus *int32 `json:"foreign_switch_status,omitempty"`

	// UDP协议禁用 0-关闭 1-开启
	UdpSwitchStatus *int32 `json:"udp_switch_status,omitempty"`
}

func (o Ips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ips struct{}"
	}

	return strings.Join([]string{"Ips", string(data)}, " ")
}
