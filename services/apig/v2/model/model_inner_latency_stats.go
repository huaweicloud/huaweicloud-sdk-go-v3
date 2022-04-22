package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InnerLatencyStats struct {

	// 最大网关内部延时  单位：ms
	MaxInnerLatency *int32 `json:"max_inner_latency,omitempty"`

	// 平均网关内部延时  单位：ms
	AvgInnerLatency *float32 `json:"avg_inner_latency,omitempty"`
}

func (o InnerLatencyStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnerLatencyStats struct{}"
	}

	return strings.Join([]string{"InnerLatencyStats", string(data)}, " ")
}
