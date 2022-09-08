package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *CreateCommonTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
