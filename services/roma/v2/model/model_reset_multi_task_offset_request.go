package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetMultiTaskOffsetRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 任务ID

	TaskId string `json:"task_id"`

	Body *MultiTaskResetBody `json:"body,omitempty"`
}

func (o ResetMultiTaskOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMultiTaskOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetMultiTaskOffsetRequest", string(data)}, " ")
}
