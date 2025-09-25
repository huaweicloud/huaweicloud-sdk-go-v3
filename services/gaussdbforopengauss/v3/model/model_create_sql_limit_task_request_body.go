package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlLimitTaskRequestBody struct {

	// **参数解释**: 限流任务范围，目前支持SQL、SESSION级别范围。 **约束限制**: 不涉及。 **取值范围**: - SQL - SESSION  **默认取值**: 不涉及。
	TaskScope string `json:"task_scope"`

	// **参数解释**: 任务开始时间，如果该值小于当前时间，会取当前时间的前两分钟。 **约束限制**: 当“task_scope”为SQL时必传。 **取值范围**: 格式必须为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。 **默认取值**: 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 任务结束时间。 **约束限制**: 当“task_scope”为SQL时必传。 **取值范围**: 大于任务开始时间，格式必须为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。 **默认取值**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 限流类型。 **约束限制**: 不涉及。 **取值范围**: - 当“task_scope”为SQL时，可选SQL_ID、SQL_TYPE类型。 - 当“task_scope”为SESSION时，可选SESSION_ACTIVE_MAX_COUNT类型。  **默认取值**: 不涉及。
	LimitType string `json:"limit_type"`

	// **参数解释**: 限流类型值。 **约束限制**: 不涉及。 **取值范围**: - 当“limit_type”为SQL_ID类型时，该值为选中模板的sql_id。 - 当“limit_type”为SQL_TYPE类型时，值为SQL类型，目前支持select，update，insert，delete，merge。 - 当“limit_type”为SESSION_ACTIVE_MAX_COUNT类型时，只支持CPU_OR_MEMORY一种值。  **默认取值**: 不涉及。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// **参数解释**: 关键词。 **约束限制**: 当且仅当“limit_type”为SQL_TYPE时，必传。 **取值范围**: 多个关键词以逗号隔开，数量范围为[2，100]个，每个关键词长度范围[2，64]位，关键词不允许包含 \" 或 \\ 或 {} 或 null值以及非首尾的空格符。 **默认取值**: 不涉及。
	KeyWords *string `json:"key_words,omitempty"`

	// **参数解释**: 限流任务名。 **约束限制**: 不涉及。 **取值范围**: 只能为英文字母大小写，下划线，数字和$符，最大长度为100个字符。 **默认取值**: 不涉及。
	TaskName string `json:"task_name"`

	// **参数解释**: 并发数。 **约束限制**: 不涉及。 **取值范围**: 大于等于零的整数，取值范围[0, 2147483647]。 **默认取值**: 不涉及。
	ParallelSize int32 `json:"parallel_size"`

	// **参数解释**: CPU利用率阈值。 **约束限制**: 如果“limit_type”为SESSION_ACTIVE_MAX_COUNT，必传。不支持和内存利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。 **取值范围**: 整数，取值范围[0,100）。 **默认取值**: 不涉及。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// **参数解释**: 内存利用率阈值。 **约束限制**: 如果“limit_type”为SESSION_ACTIVE_MAX_COUNT，必传。不支持和CPU利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。 **取值范围**: 整数，取值范围[0,100）。 **默认取值**: 不涉及。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// **参数解释**: 数据库名称，限流只对指定的数据库生效，多个数据库名称用英文逗号分割。 **约束限制**: 当“limit_type”为SQL_TYPE类型时，databases参数必选。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Databases *string `json:"databases,omitempty"`

	// **参数解释**: CN节点信息列表 **约束限制**: 如果“limit_type”为SQL_ID，则“node_infos”必选。
	NodeInfos *[]CreateLimitTaskNodeOption `json:"node_infos,omitempty"`
}

func (o CreateSqlLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitTaskRequestBody", string(data)}, " ")
}
