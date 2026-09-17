package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineRunsQuery 查询流水线运行历史请求体
type ListPipelineRunsQuery struct {

	// **参数解释**： 流水线状态列表。 **约束限制**： 不涉及。 **取值范围**： - COMPLETED：已完成。 - RUNNING：运行中。 - FAILED：失败。 - CANCELED：取消。 - PAUSED：暂停。 - SUSPEND：挂起。 - IGNORED：忽略。 **默认取值**： 不涉及。
	Status *[]string `json:"status,omitempty"`

	// **参数解释**： 流水线开始时间。 **约束限制**： 不涉及。 **取值范围**： 时间戳或者yyyy-MM-dd HH:mm:ss格式均可。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 流水线结束时间。 **约束限制**： 不涉及。 **取值范围**： 时间戳或者yyyy-MM-dd HH:mm:ss格式均可。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 流水线状态更新时间。 **约束限制**： 不涉及。 **取值范围**： 时间戳或者yyyy-MM-dd HH:mm:ss格式均可。 **默认取值**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 触发类型列表。 **约束限制**： 不涉及。 **取值范围**： - Manual：手动触发。 - Scheduler：定时触发。 - RollBack：回退触发。 - CreateTag：Tag事件触发。 - Note：评论触发。 - Issue：Issue触发。 - MR：MR触发。 - CR：CR触发。 - Generic：流水线触发器触发。 - Push：Push事件触发。 - SubPipeline：子流水线触发。 **默认取值**： 不涉及。
	TriggerType *[]string `json:"trigger_type,omitempty"`

	// **参数解释**： 执行人ID列表。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	ExecutorIds *[]string `json:"executor_ids,omitempty"`

	// **参数解释**： 起始偏移。 **约束限制**： 不涉及。 **取值范围**： 大于等于零。 **默认取值**： 不涉及。
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 大于等于零。 **默认取值**： 不涉及。
	Limit *int64 `json:"limit,omitempty"`

	// **参数解释**： 排序字段名称。 **约束限制**： 不涉及。 **取值范围**： \"start_time\" - 流水线开始时间。 \"update_time\" - 流水线更新时间。 **默认取值**： 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及。 **取值范围**： - asc：按排序字段升序。 - desc：按排序字段降序。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 是否返回Job状态详情。 **约束限制**： 不涉及。 **取值范围**： - true：返回Job状态列表。 - false：不返回。 **默认取值**： false。
	ShowJobDetails *bool `json:"show_job_details,omitempty"`

	// **参数解释**： 阶段ID，用于指定返回Job状态详情的阶段。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及，为空时默认取流水线最后一个阶段。
	StageId *string `json:"stage_id,omitempty"`

	// **参数解释**： Job ID，仅在show_job_details为true时生效，用于过滤包含指定Job的执行记录。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`
}

func (o ListPipelineRunsQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsQuery", string(data)}, " ")
}
