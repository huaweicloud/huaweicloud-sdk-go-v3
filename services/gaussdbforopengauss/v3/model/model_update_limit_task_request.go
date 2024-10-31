package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLimitTaskRequest Request Object
type UpdateLimitTaskRequest struct {

	// 限流任务id。
	TaskId string `json:"task_id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateLimitTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateLimitTaskRequest", string(data)}, " ")
}
