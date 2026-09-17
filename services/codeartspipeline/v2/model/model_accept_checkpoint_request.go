package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptCheckpointRequest Request Object
type AcceptCheckpointRequest struct {

	// 流水线步骤ID
	StepRunId string `json:"step_run_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线运行实例ID
	PipelineRunId string `json:"pipeline_run_id"`
}

func (o AcceptCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptCheckpointRequest struct{}"
	}

	return strings.Join([]string{"AcceptCheckpointRequest", string(data)}, " ")
}
