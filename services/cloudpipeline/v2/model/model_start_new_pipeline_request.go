package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartNewPipelineRequest struct {

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	Body *StartPipelineParameters `json:"body,omitempty"`
}

func (o StartNewPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNewPipelineRequest struct{}"
	}

	return strings.Join([]string{"StartNewPipelineRequest", string(data)}, " ")
}
