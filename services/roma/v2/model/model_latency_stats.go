package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LatencyStats struct {
	// 最大延时

	MaxLatency *int32 `json:"max_latency,omitempty"`
	// 平均延时

	AvgLatency float32 `json:"avg_latency,omitempty"`
}

func (o LatencyStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatencyStats struct{}"
	}

	return strings.Join([]string{"LatencyStats", string(data)}, " ")
}
