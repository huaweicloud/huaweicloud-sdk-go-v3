package model

import (
	"encoding/json"

	"strings"
)

type InnerLatencyStats struct {
	// 最大网关内部延时  单位：ms

	MaxInnerLatency *int32 `json:"max_inner_latency,omitempty"`
	// 平均网关内部延时  单位：ms

	AvgInnerLatency float32 `json:"avg_inner_latency,omitempty"`
}

func (o InnerLatencyStats) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InnerLatencyStats struct{}"
	}

	return strings.Join([]string{"InnerLatencyStats", string(data)}, " ")
}
