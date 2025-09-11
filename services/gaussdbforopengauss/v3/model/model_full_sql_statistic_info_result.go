package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlStatisticInfoResult 全量SQL聚合统计信息。
type FullSqlStatisticInfoResult struct {

	// **参数解释**: SQL模板。未开启实例内核GUC参数（track_stmt_parameter）时，展示的是归一化SQL模板内容；开启该参数后，展示的是随机一条SQL记录中截断参数列表后的内容。 **取值范围**: 不涉及。
	Template *string `json:"template,omitempty"`

	// **参数解释**: 归一化SQL ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 汇总SQL条目数。 **取值范围**: 不涉及。
	SqlCount *int64 `json:"sql_count,omitempty"`

	// **参数解释**: 平均有效DB耗时（微秒）。 **取值范围**: 不涉及。
	AvgDbTime *int64 `json:"avg_db_time,omitempty"`

	// **参数解释**: 平均CPU执行耗时（微秒）。 **取值范围**: 不涉及。
	AvgCpuTime *int64 `json:"avg_cpu_time,omitempty"`

	// **参数解释**: 平均IO耗时（微秒）。 **取值范围**: 不涉及。
	AvgDataIoTime *int64 `json:"avg_data_io_time,omitempty"`

	// **参数解释**: 平均SQL执行器内执行时间（微秒）。 **取值范围**: 不涉及。
	AvgExecutionTime *int64 `json:"avg_execution_time,omitempty"`

	// **参数解释**: 平均Buffer块命中次数。 **取值范围**: 不涉及。
	AvgNBlocksHit *int64 `json:"avg_n_blocks_hit,omitempty"`

	// **参数解释**: 开始时间戳。 **取值范围**: 不涉及。
	StartTimeStamp *string `json:"start_time_stamp,omitempty"`

	// **参数解释**: 结束时间戳。 **取值范围**: 不涉及。
	EndTimeStamp *string `json:"end_time_stamp,omitempty"`
}

func (o FullSqlStatisticInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlStatisticInfoResult struct{}"
	}

	return strings.Join([]string{"FullSqlStatisticInfoResult", string(data)}, " ")
}
