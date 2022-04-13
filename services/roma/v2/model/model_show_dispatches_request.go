package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDispatchesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o ShowDispatchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDispatchesRequest struct{}"
	}

	return strings.Join([]string{"ShowDispatchesRequest", string(data)}, " ")
}
