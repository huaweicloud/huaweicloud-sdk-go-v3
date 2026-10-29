package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IncConfigV2 **参数解释：** 基于某一任务创建分支任务、MR任务时使用的参数。 **约束限制：** 不涉及。
type IncConfigV2 struct {

	// **参数解释：** 需要关联的任务ID。 **约束限制：** 32位UUID。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ParentTaskId *string `json:"parent_task_id,omitempty"`

	// **参数解释：** MR增量任务的源分支。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GitSourceBranch *string `json:"git_source_branch,omitempty"`

	// **参数解释：** MR增量任务的目标分支。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GitTargetBranch *string `json:"git_target_branch,omitempty"`

	// **参数解释：** MR编号。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	MergeId *string `json:"merge_id,omitempty"`

	// **参数解释：** 创建MR增量任务时需传。 **约束限制：** 不涉及。 **取值范围：** merge_request。 **默认取值：** 不涉及。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释：** 创建MR增量任务时需传。 **约束限制：** 不涉及。 **取值范围：** open。 **默认取值：** 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释：** 创建MR增量任务时需传，MR标题。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Title *string `json:"title,omitempty"`
}

func (o IncConfigV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncConfigV2 struct{}"
	}

	return strings.Join([]string{"IncConfigV2", string(data)}, " ")
}
