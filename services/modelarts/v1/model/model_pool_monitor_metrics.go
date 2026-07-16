package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolMonitorMetrics struct {
	Metric *PoolMonitorMetric `json:"metric,omitempty"`

	// **参数解释**：监控指标数据。
	DataPoints *[]PoolMonitorDataPoints `json:"dataPoints,omitempty"`
}

func (o PoolMonitorMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMonitorMetrics struct{}"
	}

	return strings.Join([]string{"PoolMonitorMetrics", string(data)}, " ")
}
