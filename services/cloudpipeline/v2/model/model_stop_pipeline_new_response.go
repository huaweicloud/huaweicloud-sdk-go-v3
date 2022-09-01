package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopPipelineNewResponse struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 流水线名字
	PipelineName   *string `json:"pipeline_name,omitempty" xml:"pipeline_name"`
	HttpStatusCode int     `json:"-"`
}

func (o StopPipelineNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineNewResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineNewResponse", string(data)}, " ")
}
