package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTaskRequest struct {
	// 服务名称

	ServiceName string `json:"service_name"`
	// 指定的服务作业ID

	TaskId string `json:"task_id"`
}

func (o DeleteTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteTaskRequest", string(data)}, " ")
}
