package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LatencyStats struct {

	// 最大延时
	MaxLatency *int32 `json:"max_latency,omitempty" xml:"max_latency"`

	// 平均延时
	AvgLatency *float32 `json:"avg_latency,omitempty" xml:"avg_latency"`
}

func (o LatencyStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatencyStats struct{}"
	}

	return strings.Join([]string{"LatencyStats", string(data)}, " ")
}
