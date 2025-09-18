package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineByTemplateIdResponse Response Object
type CreatePipelineByTemplateIdResponse struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineByTemplateIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateIdResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateIdResponse", string(data)}, " ")
}
