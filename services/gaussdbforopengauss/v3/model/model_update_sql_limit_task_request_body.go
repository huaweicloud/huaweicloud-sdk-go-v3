package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSqlLimitTaskRequestBody struct {

	// **参数解释**: 任务开始时间，如果该值小于当前时间，会取当前时间的前两分钟。 **约束限制**: 当“task_scope”为SQL时必传。 **取值范围**: 格式必须为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。 **默认取值**: 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 任务结束时间。 **约束限制**: 当“task_scope”为SQL时必传。 **取值范围**: 格式为yyyy-mm-ddThh:mm:ssZ，大于任务开始时间，当前时间指UTC时间。 **默认取值**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 关键词。 **约束限制**: 当且仅当“limit_type”为SQL_TYPE，选传。 **取值范围**: 多个关键词以逗号隔开，最大长度为64位，最大关键词数量为100个。 **默认取值**: 不涉及。
	KeyWords *string `json:"key_words,omitempty"`

	// **参数解释**: 并发数。 **约束限制**: 不涉及。 **取值范围**: 大于等于零的整数。 **默认取值**: 不涉及。
	ParallelSize int32 `json:"parallel_size"`

	// **参数解释**: 限流任务名。 **约束限制**: 不涉及。 **取值范围**: 只能为英文字母大小写，下划线，数字和$符。 **默认取值**: 不涉及。
	TaskName string `json:"task_name"`

	// **参数解释**: CPU利用率阈值。 **约束限制**: 如果“limit_type”为SESSION_ACTIVE_MAX_COUNT，与内存利用率两者至少传一个。 **取值范围**: 整数，取值范围[0,100）。 **默认取值**: 不涉及。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// **参数解释**: 内存利用率阈值。 **约束限制**: 如果“limit_type”为SESSION_ACTIVE_MAX_COUNT，与CPU利用率两者至少传一个。 **取值范围**: 整数，取值范围[0,100）。 **默认取值**: 不涉及。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// **参数解释**: CN节点数据库组。 **约束限制**: 不涉及。 **取值范围**: 每个数据库字符串以逗号形式隔开。 **默认取值**: 不涉及。
	Databases *string `json:"databases,omitempty"`

	// **参数解释**: 限流任务所在的节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`
}

func (o UpdateSqlLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitTaskRequestBody", string(data)}, " ")
}
