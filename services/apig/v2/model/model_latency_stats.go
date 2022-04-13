package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LatencyStats struct {
	// 最大延时  单位：ms

	MaxLatency *int32 `json:"max_latency,omitempty"`
	// 平均延时  单位：ms

	AvgLatency *float32 `json:"avg_latency,omitempty"`
}

func (o LatencyStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatencyStats struct{}"
	}

	return strings.Join([]string{"LatencyStats", string(data)}, " ")
}
