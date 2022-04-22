package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMultiTaskMappingsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *MultiTaskMappingCreateBody `json:"body,omitempty"`
}

func (o CreateMultiTaskMappingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiTaskMappingsRequest struct{}"
	}

	return strings.Join([]string{"CreateMultiTaskMappingsRequest", string(data)}, " ")
}
