package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStepOutputsRequest Request Object
type ShowStepOutputsRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线运行实例ID
	PipelineRunId string `json:"pipeline_run_id"`

	StepRunIds *string `json:"step_run_ids,omitempty"`
}

func (o ShowStepOutputsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStepOutputsRequest struct{}"
	}

	return strings.Join([]string{"ShowStepOutputsRequest", string(data)}, " ")
}
