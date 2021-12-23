package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ddos状态响应体
type DDosStatus struct {
	// EIP的ID

	FloatingIpId string `json:"floating_ip_id"`
	// 浮动IP地址

	FloatingIpAddress string `json:"floating_ip_address"`
	// EIP所属类型，可选范围： - EIP：未绑定到ECS的EIP或绑定到ECS的EIP - ELB：绑定到ELB的EIP

	NetworkType string `json:"network_type"`
	// 防护状态，可选范围： - normal：表示正常 - configging：表示设置中 - notConfig：表示未设置 - packetcleaning：表示清洗 - packetdropping：表示黑洞

	Status string `json:"status"`
	// 黑洞结束时间

	BlackholeEndtime int64 `json:"blackhole_endtime"`
	// 防护类型

	ProtectType string `json:"protect_type"`
	// 流量阈值

	TrafficThreshold int64 `json:"traffic_threshold"`
	// http流量阈值

	HttpThreshold int64 `json:"http_threshold"`
}

func (o DDosStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DDosStatus struct{}"
	}

	return strings.Join([]string{"DDosStatus", string(data)}, " ")
}
