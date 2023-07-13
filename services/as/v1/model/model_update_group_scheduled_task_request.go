package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupScheduledTaskRequest Request Object
type UpdateGroupScheduledTaskRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	// 计划任务ID
	ScheduledTaskId string `json:"scheduled_task_id"`

	Body *UpdateScheduledTaskOption `json:"body,omitempty"`
}

func (o UpdateGroupScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupScheduledTaskRequest", string(data)}, " ")
}
