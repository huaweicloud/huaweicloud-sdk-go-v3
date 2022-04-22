package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTaskRequest struct {

	// 服务名称
	ServiceName string `json:"service_name"`

	// 指定的服务作业ID
	TaskId string `json:"task_id"`

	Body *UpdateTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
