package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlTraceStatisticsResult 链路详情信息
type FullSqlTraceStatisticsResult struct {

	// **参数解释**: 命中率。 **取值范围**: 不涉及。
	HitRate *float64 `json:"hit_rate,omitempty"`

	// **参数解释**: 有效DB时间花费。 **取值范围**: 不涉及。
	DbTime *int32 `json:"db_time,omitempty"`

	// **参数解释**: CPU时间（单位：微秒）。 **取值范围**: 不涉及。
	CpuTime *int32 `json:"cpu_time,omitempty"`

	// **参数解释**: IO时间（单位：微秒）。 **取值范围**: 不涉及。
	IoTime *int32 `json:"io_time,omitempty"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	ExecutionTime *int32 `json:"execution_time,omitempty"`

	// **参数解释**: 扫描行数。 **取值范围**: 不涉及。
	ScanRows *int32 `json:"scan_rows,omitempty"`

	// **参数解释**: 更新行数。 **取值范围**: 不涉及。
	UpdateRows *int32 `json:"update_rows,omitempty"`

	// **参数解释**: 插入行数。 **取值范围**: 不涉及。
	InsertRows *int32 `json:"insert_rows,omitempty"`

	// **参数解释**: 删除行数。 **取值范围**: 不涉及。
	DeleteRows *int32 `json:"delete_rows,omitempty"`
}

func (o FullSqlTraceStatisticsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlTraceStatisticsResult struct{}"
	}

	return strings.Join([]string{"FullSqlTraceStatisticsResult", string(data)}, " ")
}
