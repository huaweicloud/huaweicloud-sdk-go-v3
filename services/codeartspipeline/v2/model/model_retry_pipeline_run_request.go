package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryPipelineRunRequest Request Object
type RetryPipelineRunRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线运行实例ID
	PipelineRunId string `json:"pipeline_run_id"`
}

func (o RetryPipelineRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryPipelineRunRequest struct{}"
	}

	return strings.Join([]string{"RetryPipelineRunRequest", string(data)}, " ")
}
