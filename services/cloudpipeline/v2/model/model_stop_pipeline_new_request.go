package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPipelineNewRequest Request Object
type StopPipelineNewRequest struct {

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线执行ID
	BuildId string `json:"build_id"`
}

func (o StopPipelineNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineNewRequest struct{}"
	}

	return strings.Join([]string{"StopPipelineNewRequest", string(data)}, " ")
}
