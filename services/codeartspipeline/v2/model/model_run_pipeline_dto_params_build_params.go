package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineDtoParamsBuildParams **参数解释**： 流水线具体构建参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RunPipelineDtoParamsBuildParams struct {

	// **参数解释**： 代码仓触发类型，包含branch-分支触发，tag-标签触发等。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BuildType string `json:"build_type"`

	// **参数解释**： 流水线的触发方式。 - Manual：手动触发。 - Scheduler：定时任务。 - MR：MR触发。 - Push：Push事件触发。 - CreateTag：Tag事件触发。 - Issue：Issue触发。 - Note：评论触发。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释**： 流水线触发运行分支。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释**： 流水线触发运行的标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tag *string `json:"tag,omitempty"`
}

func (o RunPipelineDtoParamsBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoParamsBuildParams struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoParamsBuildParams", string(data)}, " ")
}
