package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlInfoResult struct {

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: SCHEMA名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 变量替换后的完整SQL。当sql_text不返回变量值时，sql返回空字符串。 **取值范围**: 不涉及。
	Sql *string `json:"sql,omitempty"`

	// **参数解释**: SQL ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 用户名称。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: SQL文本。 **取值范围**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: SQL计划。 **取值范围**: 不涉及。
	QueryPlan *string `json:"query_plan,omitempty"`

	// **参数解释**: 执行次数（次）。 **取值范围**: 不涉及。
	Calls *int32 `json:"calls,omitempty"`

	// **参数解释**: 平均执行时间（us）。 **取值范围**: 不涉及。
	AvgExecTime *string `json:"avg_exec_time,omitempty"`

	// **参数解释**: 平均CPU耗时（us）。 **取值范围**: 不涉及。
	AvgCpuTime *string `json:"avg_cpu_time,omitempty"`

	// **参数解释**: 平均IO耗时（us）。 **取值范围**: 不涉及。
	AvgIoTime *string `json:"avg_io_time,omitempty"`

	// **参数解释**: 平均返回行数（行）。 **取值范围**: 不涉及。
	AvgReturnedRows *int32 `json:"avg_returned_rows,omitempty"`

	// **参数解释**: 平均扫描行数（行）。 **取值范围**: 不涉及。
	AvgFetchedRows *int32 `json:"avg_fetched_rows,omitempty"`

	// **参数解释**: buffer命中率。 **取值范围**: 不涉及。
	BufferHitRatio *string `json:"buffer_hit_ratio,omitempty"`

	// **参数解释**: SQL命中率。 **取值范围**: 不涉及。
	SqlHitRatio *string `json:"sql_hit_ratio,omitempty"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`
}

func (o SlowSqlInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlInfoResult struct{}"
	}

	return strings.Join([]string{"SlowSqlInfoResult", string(data)}, " ")
}
