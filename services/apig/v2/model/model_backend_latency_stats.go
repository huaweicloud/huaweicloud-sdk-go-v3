package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackendLatencyStats struct {

	// 最大后端延时
	MaxBackendLatency *int32 `json:"max_backend_latency,omitempty" xml:"max_backend_latency"`

	// 平均后端延时
	AvgBackendLatency *float32 `json:"avg_backend_latency,omitempty" xml:"avg_backend_latency"`
}

func (o BackendLatencyStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendLatencyStats struct{}"
	}

	return strings.Join([]string{"BackendLatencyStats", string(data)}, " ")
}
