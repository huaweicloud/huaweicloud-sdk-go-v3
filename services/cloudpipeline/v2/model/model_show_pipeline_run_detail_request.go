package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineRunDetailRequest Request Object
type ShowPipelineRunDetailRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线运行实例ID
	PipelineRunId *string `json:"pipeline_run_id,omitempty"`
}

func (o ShowPipelineRunDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineRunDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineRunDetailRequest", string(data)}, " ")
}
