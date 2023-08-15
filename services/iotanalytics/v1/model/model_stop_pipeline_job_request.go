package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPipelineJobRequest Request Object
type StopPipelineJobRequest struct {

	// 管道ID
	PipelineId string `json:"pipeline_id"`

	// 停止管道作业触发savepoint
	TriggerSavepoint *bool `json:"trigger_savepoint,omitempty"`
}

func (o StopPipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineJobRequest struct{}"
	}

	return strings.Join([]string{"StopPipelineJobRequest", string(data)}, " ")
}
