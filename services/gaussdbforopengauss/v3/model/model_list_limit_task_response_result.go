package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLimitTaskResponseResult struct {

	// 限流任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务限流范围。
	TaskScope *string `json:"task_scope,omitempty"`

	// 任务限流类型。
	LimitType *string `json:"limit_type,omitempty"`

	// 任务限流类型值。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// 限流任务名。
	TaskName *string `json:"task_name,omitempty"`

	// CN节点数据库组,每个数据库字符串以逗号形式隔开。
	Databases *string `json:"databases,omitempty"`

	// SQL模板,仅当任务类型为SQL_ID时，返回该值。
	SqlModel *string `json:"sql_model,omitempty"`

	// 关键词，仅当任务类型为SQL_TYPE时，返回该值。
	KeyWords *string `json:"key_words,omitempty"`

	// 限流任务状态，当前支持：CREATING，UPDATEING，DELETING，WAIT_EXCUTE，EXCUTING，TIME_OVER，DELETED，CREATE_FAILED，UPDATE_FAILED，DELETE_FAILED，EXCEPTION，NODE_SHUT_DOWN。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// 规则名。
	RuleName *string `json:"rule_name,omitempty"`

	// 并发数。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// 限流任务开始时间,格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 限流任务结束时间,格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// cpu利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 内存利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// 创建时间为本地时间，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	Created *string `json:"created,omitempty"`

	// 更新时间为本地时间，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	Updated *string `json:"updated,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// CN节点信息列表。
	NodeInfos *[]ShowLimitTaskNodeOption `json:"node_infos,omitempty"`
}

func (o ListLimitTaskResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLimitTaskResponseResult struct{}"
	}

	return strings.Join([]string{"ListLimitTaskResponseResult", string(data)}, " ")
}
