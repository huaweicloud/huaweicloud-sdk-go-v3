package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TriggerBpsDict 流量限制列表
type TriggerBpsDict struct {

	// 流量分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
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
