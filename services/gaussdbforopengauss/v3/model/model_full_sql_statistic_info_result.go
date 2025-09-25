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

	// **参数解释**: 总SQL耗时（微秒）。 **取值范围**: 不涉及。
	TotalSqlTime *int64 `json:"total_sql_time,omitempty"`

	// **参数解释**: 平均SQL耗时（微秒）。 **取值范围**: 不涉及。
	AvgSqlTime *int64 `json:"avg_sql_time,omitempty"`

	// **参数解释**: 总有效DB耗时（微秒）。 **取值范围**: 不涉及。
	TotalDbTime *int64 `json:"total_db_time,omitempty"`

	// **参数解释**: 总CPU耗时（微秒）。 **取值范围**: 不涉及。
	TotalCpuTime *int64 `json:"total_cpu_time,omitempty"`

	// **参数解释**: 平均解释器时间（微秒）。 **取值范围**: 不涉及。
	AvgParseTime *int64 `json:"avg_parse_time,omitempty"`

	// **参数解释**: 平均执行计划时间（微秒）。 **取值范围**: 不涉及。
	AvgPlanTime *int64 `json:"avg_plan_time,omitempty"`

	// **参数解释**: 总IO耗时（微秒）。 **取值范围**: 不涉及。
	TotalDataIoTime *int64 `json:"total_data_io_time,omitempty"`

	// **参数解释**: 平均返回行数。 **取值范围**: 不涉及。
	AvgNReturnedRows *int64 `json:"avg_n_returned_rows,omitempty"`

	// **参数解释**: 平均扫描行数。 **取值范围**: 不涉及。
	AvgNTuplesFetched *int64 `json:"avg_n_tuples_fetched,omitempty"`

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
	StartTimeStamp *int64 `json:"start_time_stamp,omitempty"`

	// **参数解释**: 结束时间戳。 **取值范围**: 不涉及。
	EndTimeStamp *int64 `json:"end_time_stamp,omitempty"`
}

func (o FullSqlStatisticInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlStatisticInfoResult struct{}"
	}

	return strings.Join([]string{"FullSqlStatisticInfoResult", string(data)}, " ")
}
