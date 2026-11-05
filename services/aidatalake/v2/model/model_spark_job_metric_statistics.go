package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobMetricStatistics **参数解释**：SparkJob运行实例上资源的指标信息。
type SparkJobMetricStatistics struct {

	// **参数解释**：当前CPU使用量，单位毫核。 **取值范围**：大于等于1000。
	CurrentCpu *int64 `json:"current_cpu,omitempty"`

	// **参数解释**：当前内存使用量，单位MB。 **取值范围**：大于等于1024。
	CurrentMemory *int64 `json:"current_memory,omitempty"`

	// **参数解释**：估算 CU 秒（折算后的秒数，四舍五入取整，如 262）。
	TotalCuSeconds *int64 `json:"total_cu_seconds,omitempty"`
}

func (o SparkJobMetricStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobMetricStatistics struct{}"
	}

	return strings.Join([]string{"SparkJobMetricStatistics", string(data)}, " ")
}
