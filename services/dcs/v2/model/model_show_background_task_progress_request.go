package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundTaskProgressRequest Request Object
type ShowBackgroundTaskProgressRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 任务ID。
	TaskId string `json:"task_id"`
}

func (o ShowBackgroundTaskProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskProgressRequest", string(data)}, " ")
}
