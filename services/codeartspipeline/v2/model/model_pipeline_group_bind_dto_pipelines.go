package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineGroupBindDtoPipelines struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	PipelineId string `json:"pipeline_id"`

	// **参数解释**： 流水线名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PipelineName string `json:"pipeline_name"`
}

func (o PipelineGroupBindDtoPipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupBindDtoPipelines struct{}"
	}

	return strings.Join([]string{"PipelineGroupBindDtoPipelines", string(data)}, " ")
}
