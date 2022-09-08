package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 扩展配置列表
type ExtendDDoSSet struct {

	// 配置分段ID
	SetID int64 `json:"SetID"`

	// 单一源IP新建连接个数
	NewConnectionLimited int64 `json:"new_connection_limited"`

	// 单一源IP连接数总个数
	TotalConnectionLimited int64 `json:"total_connection_limited"`

	// 每秒HTTP请求数（个/s）阈值
	HttpPacketPerSecond int64 `json:"http_packet_per_second"`

	// 每秒流量（Mbit/s）阈值
	TrafficPerSecond int64 `json:"traffic_per_second"`

	// 每秒报文数（个/s）阈值
	PacketPerSecond int64 `json:"packet_per_second"`
}

func (o ExtendDDoSSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendDDoSSet struct{}"
	}

	return strings.Join([]string{"ExtendDDoSSet", string(data)}, " ")
}
