package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLimitTaskRequest Request Object
type DeleteLimitTaskRequest struct {

	// 限流任务id。
	TaskId string `json:"task_id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteLimitTaskRequest", string(data)}, " ")
}
