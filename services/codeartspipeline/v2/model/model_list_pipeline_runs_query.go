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

	// **参数解释**： 起始偏移。 **约束限制**： 不涉及。 **取值范围**： 大于等于零。 **默认取值**： 不涉及。
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 大于等于零。 **默认取值**： 不涉及。
	Limit *int64 `json:"limit,omitempty"`

	// **参数解释**： 排序字段名称。 **约束限制**： 不涉及。 **取值范围**： \"start_time\" - 流水线开始时间。 **默认取值**： 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及。 **取值范围**： - asc：按排序字段升序。 - desc：按排序字段降序。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListPipelineRunsQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsQuery", string(data)}, " ")
}
