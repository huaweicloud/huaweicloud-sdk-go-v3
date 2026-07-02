package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaInfo struct {

	// 网段协商ID
	Id *string `json:"id,omitempty"`

	// 源IP网段
	SourceIpCidr *string `json:"source_ip_cidr,omitempty"`

	// 目的IP网段
	DestIpCidr *string `json:"dest_ip_cidr,omitempty"`

	// 发送包
	PacketsSent float32 `json:"packets_sent,omitempty"`

	// 接收包
	PacketsRecv float32 `json:"packets_recv,omitempty"`

	// 发送流量(Byte)
	TrafficSent float32 `json:"traffic_sent,omitempty"`

	// 接收流量(Byte)
	TrafficRecv float32 `json:"traffic_recv,omitempty"`

	// 数据收集时间
	CollectedAt *sdktime.SdkTime `json:"collected_at,omitempty"`
}

func (o SaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaInfo struct{}"
	}

	return strings.Join([]string{"SaInfo", string(data)}, " ")
}
