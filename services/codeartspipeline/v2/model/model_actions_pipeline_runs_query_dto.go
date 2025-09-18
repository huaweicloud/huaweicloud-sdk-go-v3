package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionsPipelineRunsQueryDto struct {

	// **参数解释**： 分页查询页码。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 不涉及。
	Page *int64 `json:"page,omitempty"`

	// **参数解释**： 每页的查询数量。 **约束限制**： 不涉及。 **取值范围**： 固定20。 **默认取值**： 不涉及。
	PageSize *string `json:"page_size,omitempty"`

	// **参数解释**： 代码源地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HttpsUrl string `json:"https_url"`

	// **参数解释**： 流水线名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PipelineName *string `json:"pipeline_name,omitempty"`

	// **参数解释**： 文件所处路径。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 流水线运行人名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PipelineRunName *string `json:"pipeline_run_name,omitempty"`

	// **参数解释**： 流水线触发类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Event *string `json:"event,omitempty"`

	// **参数解释**： 流水线创建者名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Actor *string `json:"actor,omitempty"`

	// **参数解释**： 代码源分支。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Branch *string `json:"branch,omitempty"`

	// **参数解释**： 流水线运行状态。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`
}

func (o ActionsPipelineRunsQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionsPipelineRunsQueryDto struct{}"
	}

	return strings.Join([]string{"ActionsPipelineRunsQueryDto", string(data)}, " ")
}
