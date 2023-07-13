package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupScheduledTaskRequest Request Object
type DeleteGroupScheduledTaskRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	// 计划任务ID
	ScheduledTaskId string `json:"scheduled_task_id"`
}

func (o DeleteGroupScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupScheduledTaskRequest", string(data)}, " ")
}
