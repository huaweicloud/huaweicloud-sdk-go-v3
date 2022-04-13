package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopPipelineNewResponse struct {
	// 流水线ID

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名字

	PipelineName   *string `json:"pipeline_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopPipelineNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineNewResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineNewResponse", string(data)}, " ")
}
