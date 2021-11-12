package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemovePipelineResponse struct {
	// 流水线ID

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名字

	PipelineName   *string `json:"pipeline_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemovePipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePipelineResponse struct{}"
	}

	return strings.Join([]string{"RemovePipelineResponse", string(data)}, " ")
}
