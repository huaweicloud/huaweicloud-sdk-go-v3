package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlDetailResult struct {

	// **参数解释**: SQL ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 用户名称。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 客户端IP。 **取值范围**: 不涉及。
	ClientIp *string `json:"client_ip,omitempty"`

	// **参数解释**: 客户端端口。 **取值范围**: 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**: SQL模版。 **取值范围**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: 执行计划。 **取值范围**: 不涉及。
	QueryPlan *string `json:"query_plan,omitempty"`

	// **参数解释**: 开始时间UTC时间。 **取值范围**: 格式为yyyy-mm-ddThh:mm:ss+0000。
	StartTime *int32 `json:"start_time,omitempty"`

	// **参数解释**: 结束时间UTC时间。 **取值范围**: 格式为yyyy-mm-ddThh:mm:ss+0000。
	FinishTime *int32 `json:"finish_time,omitempty"`

	// **参数解释**: 返回行。 **取值范围**: 不涉及。
	ReturnedRows *int32 `json:"returned_rows,omitempty"`

	// **参数解释**: 扫描行。 **取值范围**: 不涉及。
	FetchedRows *int32 `json:"fetched_rows,omitempty"`

	// **参数解释**: 扫描页。 **取值范围**: 不涉及。
	FetchedPages *int32 `json:"fetched_pages,omitempty"`

	// **参数解释**: 命中页。 **取值范围**: 不涉及。
	HitPages *int32 `json:"hit_pages,omitempty"`

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	TotalTime *int32 `json:"total_time,omitempty"`

	// **参数解释**: CPU耗时（单位：微秒）。 **取值范围**: 不涉及。
	CpuTime *int32 `json:"cpu_time,omitempty"`

	// **参数解释**: 计划耗时（单位：微秒）。 **取值范围**: 不涉及。
	PlanTime *int32 `json:"plan_time,omitempty"`

	// **参数解释**: IO耗时（单位：微秒）。 **取值范围**: 不涉及。
	IoTime *int32 `json:"io_time,omitempty"`

	// **参数解释**: 加锁次数。 **取值范围**: 不涉及。
	LockCount *int32 `json:"lock_count,omitempty"`

	// **参数解释**: 加锁耗时(单位：微秒)。 **取值范围**: 不涉及。
	LockTime *int32 `json:"lock_time,omitempty"`
}

func (o SlowSqlDetailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlDetailResult struct{}"
	}

	return strings.Join([]string{"SlowSqlDetailResult", string(data)}, " ")
}
