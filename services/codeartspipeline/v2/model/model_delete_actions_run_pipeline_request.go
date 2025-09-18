package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActionsRunPipelineRequest Request Object
type DeleteActionsRunPipelineRequest struct {

	// **参数解释**： 租户id。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	DomainId string `json:"domain_id"`

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	PipelineId string `json:"pipeline_id"`

	// **参数解释**： 流水线运行实例ID，[运行流水线](RunPipeline.xml)接口的返回值即为流水线运行实例ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	PipelineRunId string `json:"pipeline_run_id"`
}

func (o DeleteActionsRunPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActionsRunPipelineRequest struct{}"
	}

	return strings.Join([]string{"DeleteActionsRunPipelineRequest", string(data)}, " ")
}
