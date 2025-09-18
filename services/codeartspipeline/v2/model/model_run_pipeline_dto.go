package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineDto 运行流水线请求体
type RunPipelineDto struct {

	// **参数解释**： 代码源信息列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sources *[]RunPipelineDtoSources `json:"sources,omitempty"`

	// **参数解释**： 流水线运行描述。 **约束限制**： 不涉及。 **取值范围**： 不超过1024字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 使用的自定义参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables *[]RunPipelineDtoVariables `json:"variables,omitempty"`

	// **参数解释**： 流水线运行时选择的流水线任务。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ChooseJobs *[]string `json:"choose_jobs,omitempty"`

	// **参数解释**： 选择的流水线阶段。优先级高于choose_jobs，即stage未选择时，无视choose_jobs中该stage下的job。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ChooseStages *[]string `json:"choose_stages,omitempty"`
}

func (o RunPipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDto struct{}"
	}

	return strings.Join([]string{"RunPipelineDto", string(data)}, " ")
}
