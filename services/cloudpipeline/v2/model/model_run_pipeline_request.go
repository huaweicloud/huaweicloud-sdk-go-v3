package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineRequest Request Object
type RunPipelineRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	Body *RunPipelineDto `json:"body,omitempty"`
}

func (o RunPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineRequest struct{}"
	}

	return strings.Join([]string{"RunPipelineRequest", string(data)}, " ")
}
