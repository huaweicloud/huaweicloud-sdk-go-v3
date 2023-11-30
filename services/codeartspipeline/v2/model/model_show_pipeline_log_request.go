package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineLogRequest Request Object
type ShowPipelineLogRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线运行实例ID
	PipelineRunId string `json:"pipeline_run_id"`

	// 流水线任务ID
	JobRunId string `json:"job_run_id"`

	// 流水线步骤ID
	StepRunId string `json:"step_run_id"`

	Body *LogQuery `json:"body,omitempty"`
}

func (o ShowPipelineLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLogRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineLogRequest", string(data)}, " ")
}
