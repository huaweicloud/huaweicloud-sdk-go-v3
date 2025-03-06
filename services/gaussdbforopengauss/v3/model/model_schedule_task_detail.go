package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleTaskDetail struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务创建时间,格式为yyyy-mm-ddThh:mm:ssZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务开始时间,格式为yyyy-mm-ddThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 任务信息。
	TaskContent *interface{} `json:"task_content,omitempty"`
}

func (o ScheduleTaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTaskDetail struct{}"
	}

	return strings.Join([]string{"ScheduleTaskDetail", string(data)}, " ")
}
