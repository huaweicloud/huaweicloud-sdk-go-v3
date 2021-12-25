package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTaskRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o RunTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskRequest struct{}"
	}

	return strings.Join([]string{"RunTaskRequest", string(data)}, " ")
}
