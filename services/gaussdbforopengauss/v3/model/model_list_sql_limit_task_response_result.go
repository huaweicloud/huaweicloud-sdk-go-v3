package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSqlLimitTaskResponseResult struct {

	// **参数解释**: 限流任务ID。 **取值范围**: 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 任务限流范围。 **取值范围**: 不涉及。
	TaskScope *string `json:"task_scope,omitempty"`

	// **参数解释**: 任务限流类型。 **取值范围**: 不涉及。
	LimitType *string `json:"limit_type,omitempty"`

	// **参数解释**: 任务限流类型值。 **取值范围**: 不涉及。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// **参数解释**: 限流任务名。 **取值范围**: 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: CN节点数据库组。 **取值范围**: 每个数据库字符串以逗号形式隔开。
	Databases *string `json:"databases,omitempty"`

	// **参数解释**: 关键词。 **取值范围**: 不涉及。
	KeyWords *string `json:"key_words,omitempty"`

	// **参数解释**: 限流任务状态。 **取值范围**: 当前支持：CREATING，UPDATING，DELETING，WAIT_EXECUTE，EXCUTING，TIME_OVER，DELETED，CREATE_FAILED，UPDATE_FAILED，DELETE_FAILED，EXCEPTION，NODE_SHUT_DOWN。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 规则名。 **取值范围**: 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**: 并发数。 **取值范围**: 不涉及。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// **参数解释**: 限流任务开始时间。 **取值范围**: 格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 限流任务结束时间。 **取值范围**: 格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: CPU利用率。 **取值范围**: 不涉及。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// **参数解释**: 内存利用率。 **取值范围**: 不涉及。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`
}

func (o ListSqlLimitTaskResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitTaskResponseResult struct{}"
	}

	return strings.Join([]string{"ListSqlLimitTaskResponseResult", string(data)}, " ")
}
