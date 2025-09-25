package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLimitTaskResponseResult struct {

	// **参数解释**: 限流任务ID。 **取值范围**: 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 任务限流范围。 **取值范围**: 目前支持SQL，SESSION两种级别范围。
	TaskScope *string `json:"task_scope,omitempty"`

	// **参数解释**: 任务限流类型。 **取值范围**: - 当“task_scope”为SQL时，可选SQL_ID、SQL_TYPE类型。 - 当“task_scope”为SESSION时，可选SESSION_ACTIVE_MAX_COUNT类型。
	LimitType *string `json:"limit_type,omitempty"`

	// **参数解释**: 任务限流类型值。 **取值范围**: - 当“limit_type”为SQL_ID类型时，该值为选中模板的sql_id。 - 当“limit_type”为SQL_TYPE类型时，值为SQL类型，为select，update，insert，delete，merge的一种。 - 当“limit_type”为SESSION_ACTIVE_MAX_COUNT类型时，该值为CPU_OR_MEMORY。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// **参数解释**: 限流任务名。 **取值范围**: 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 实例的数据库列表，每个数据库以英文逗号形式隔开。 **取值范围**: 不涉及。
	Databases *string `json:"databases,omitempty"`

	// **参数解释**: SQL模板，仅当任务类型为SQL_ID时，返回该值。 **取值范围**: 不涉及。
	SqlModel *string `json:"sql_model,omitempty"`

	// **参数解释**: 关键词，仅当任务类型为SQL_TYPE时，返回该值。 **取值范围**: 不涉及。
	KeyWords *string `json:"key_words,omitempty"`

	// **参数解释**: 限流任务状态。 **取值范围**: 当前支持：CREATING，UPDATING，DELETING，WAIT_EXECUTE，EXCUTING，TIME_OVER，DELETED，CREATE_FAILED，UPDATE_FAILED，DELETE_FAILED，EXCEPTION，NODE_SHUT_DOWN。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 规则名。 **取值范围**: 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**: 并发数。 **取值范围**: [0, 2147483647]
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// **参数解释**: 限流任务开始时间。 **取值范围**: 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 限流任务结束时间。 **取值范围**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: CPU利用率阈值，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留整数部分。 **取值范围**: [0, 100)
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// **参数解释**: 内存利用率阈值，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留整数部分。 **取值范围**: [0, 100)
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// **参数解释**: 限流任务创建时间。 **取值范围**: 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**: 限流任务更新时间。 **取值范围**: 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**: 创建者。 **取值范围**: 不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**: 更新者。 **取值范围**: 不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释**: CN节点信息列表。
	NodeInfos *[]ShowLimitTaskNodeOption `json:"node_infos,omitempty"`
}

func (o ListLimitTaskResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLimitTaskResponseResult struct{}"
	}

	return strings.Join([]string{"ListLimitTaskResponseResult", string(data)}, " ")
}
