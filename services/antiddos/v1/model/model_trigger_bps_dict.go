package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流量限制列表
type TriggerBpsDict struct {

	// 流量分段ID
	TrafficPosId int64 `json:"traffic_pos_id"`

	// 每秒流量（Mbit/s）阈值
	TrafficPerSecond int64 `json:"traffic_per_second"`

	// 每秒报文数（个/s）阈值
	PacketPerSecond int64 `json:"packet_per_second"`
}

func (o TriggerBpsDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerBpsDict struct{}"
	}

	return strings.Join([]string{"TriggerBpsDict", string(data)}, " ")
}
