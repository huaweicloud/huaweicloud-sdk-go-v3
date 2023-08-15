package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRequest Request Object
type ShowTaskRequest struct {

	// 服务名称
	ServiceName string `json:"service_name"`

	// 指定的服务作业ID
	TaskId string `json:"task_id"`
}

func (o ShowTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
