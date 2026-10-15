package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkResourceMetric **参数解释**：Flink作业运行实例上资源的指标信息。
type FlinkResourceMetric struct {

	// **参数解释**：总内存，单位MB。 **取值范围**：0~9223372036854775807。
	TotalMemory *int64 `json:"total_memory,omitempty"`

	// **参数解释**：总CPU，单位毫核。 **取值范围**：0~9223372036854775807。
	TotalCpu *int64 `json:"total_cpu,omitempty"`

	// **参数解释**：总磁盘容量，单位GB。 **取值范围**：0~9223372036854775807。
	TotalDisk *int64 `json:"total_disk,omitempty"`
}

func (o FlinkResourceMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkResourceMetric struct{}"
	}

	return strings.Join([]string{"FlinkResourceMetric", string(data)}, " ")
}
