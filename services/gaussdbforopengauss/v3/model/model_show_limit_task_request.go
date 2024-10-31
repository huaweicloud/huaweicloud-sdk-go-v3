package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLimitTaskRequest Request Object
type ShowLimitTaskRequest struct {

	// 限流任务id。
	TaskId string `json:"task_id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowLimitTaskRequest", string(data)}, " ")
}
