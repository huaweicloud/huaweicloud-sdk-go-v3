package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTaskRequest Request Object
type StartTaskRequest struct {

	// 服务名称
	ServiceName string `json:"service_name"`

	// 指定的服务作业ID
	TaskId string `json:"task_id"`
}

func (o StartTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTaskRequest struct{}"
	}

	return strings.Join([]string{"StartTaskRequest", string(data)}, " ")
}
