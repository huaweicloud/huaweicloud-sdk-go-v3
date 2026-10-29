package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobMetricStatistics **参数解释**：SparkJob运行实例上资源的指标信息。
type SparkJobMetricStatistics struct {

	// **参数解释**：当前第几次重试，用于区分同一作业不同重试批次的指标 **取值范围**：最小值为1。
	RetryAttempt *int32 `json:"retry_attempt,omitempty"`

	// **参数解释**：指标生成时间。 **取值范围**：unix时间戳，单位为毫秒。
	MetricTime *int64 `json:"metric_time,omitempty"`

	// **参数解释**：当前CPU使用量，单位毫核。 **取值范围**：大于等于1000。
	CurrentCpu *int64 `json:"current_cpu,omitempty"`

	// **参数解释**：当前内存使用量，单位MB。 **取值范围**：大于等于1024。
	CurrentMemory *int64 `json:"current_memory,omitempty"`

	// **参数解释**：估算 CU 时（折算后的小时数，保留4位小数，如 2.3500）。 **取值范围**：大于等于0。
	TotalCuHours *float64 `json:"total_cu_hours,omitempty"`
}

func (o SparkJobMetricStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobMetricStatistics struct{}"
	}

	return strings.Join([]string{"SparkJobMetricStatistics", string(data)}, " ")
}
