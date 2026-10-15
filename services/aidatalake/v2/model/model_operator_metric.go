package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperatorMetric **参数解释**：算子监控信息。 **取值范围**：不涉及。
type OperatorMetric struct {

	// **参数解释**：查询ID。 **取值范围**：1~9223372036854775807。
	QueryId *int64 `json:"query_id,omitempty"`

	// **参数解释**：节点ID。 **取值范围**：1~2147483647。
	PlanNodeId *int32 `json:"plan_node_id,omitempty"`

	// **参数解释**：节点名称。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	PlanNodeName *string `json:"plan_node_name,omitempty"`

	// **参数解释**：开始时间，格式为：YYYY-MM-DD HH:MM:SS。 **取值范围**：时间范围：1000-01-01 00:00:00/9999-12-31 23:59:59。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// **参数解释**：执行时长。 **取值范围**：1~9223372036854775807。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：算子当前执行状态。 **取值范围**：   - init：初始化。   - waiting：等待中。   - finished：已结束。   - running：运行中。
	Status *string `json:"status,omitempty"`

	// **参数解释**：查询的并行度。 **取值范围**：1~2147483647。
	QueryDop *int32 `json:"query_dop,omitempty"`

	// **参数解释**：预估的行数。 **取值范围**：1~9223372036854775807。
	EstimatedRows *int64 `json:"estimated_rows,omitempty"`

	// **参数解释**：当前已经处理完成的行数。 **取值范围**：1~9223372036854775807。
	TupleProcessed *int64 `json:"tuple_processed,omitempty"`

	// **参数解释**：内存使用的最小峰值。 **取值范围**：1~2147483647。
	MinPeakMemory *int32 `json:"min_peak_memory,omitempty"`

	// **参数解释**：内存使用的最大峰值。 **取值范围**：1~2147483647。
	MaxPeakMemory *int32 `json:"max_peak_memory,omitempty"`

	// **参数解释**：内存使用的平均峰值。 **取值范围**：1~2147483647。
	AveragePeakMemory *int32 `json:"average_peak_memory,omitempty"`

	// **参数解释**：内存使用倾斜百分比。 **取值范围**：1~2147483647。
	MemorySkewPercent *int32 `json:"memory_skew_percent,omitempty"`

	// **参数解释**：下盘的最小数据量。 **取值范围**：1~2147483647。
	MinSpillSize *int32 `json:"min_spill_size,omitempty"`

	// **参数解释**：下盘的最大数据量。 **取值范围**：1~2147483647。
	MaxSpillSize *int32 `json:"max_spill_size,omitempty"`

	// **参数解释**：下盘的平均数据量。 **取值范围**：1~2147483647。
	AverageSpillSize *int32 `json:"average_spill_size,omitempty"`

	// **参数解释**：下盘数据的倾斜百分比。 **取值范围**：不涉及。
	SpillSkewPercent *string `json:"spill_skew_percent,omitempty"`

	// **参数解释**：算子执行迭代次数。 **取值范围**：1~2147483647。
	Loops *int32 `json:"loops,omitempty"`

	// **参数解释**：执行进度百分比。 **取值范围**：不涉及。
	Progress *string `json:"progress,omitempty"`

	// **参数解释**：所有节点读取的总数据量。 **取值范围**：1~9223372036854775807。
	TotalReadSize *int64 `json:"total_read_size,omitempty"`

	// **参数解释**：所有节点写入的总数据量。 **取值范围**：1~9223372036854775807。
	TotalWriteSize *int64 `json:"total_write_size,omitempty"`

	// **参数解释**：子节点。 **取值范围**：不涉及。
	SubOperatorMetrics *[]OperatorMetric `json:"sub_operator_metrics,omitempty"`
}

func (o OperatorMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperatorMetric struct{}"
	}

	return strings.Join([]string{"OperatorMetric", string(data)}, " ")
}
