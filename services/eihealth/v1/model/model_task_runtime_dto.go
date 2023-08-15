package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskRuntimeDto struct {

	// 作业子任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 作业子任务运行创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业子任务运行结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业子任务实际运行时间
	ActualRunningTime *int32 `json:"actual_running_time,omitempty"`

	// 作业子任务运行状态
	Status *string `json:"status,omitempty"`

	// 作业子任务的并发实例列表
	SubTasks *[]SubTaskRuntimeDto `json:"sub_tasks,omitempty"`
}

func (o TaskRuntimeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskRuntimeDto struct{}"
	}

	return strings.Join([]string{"TaskRuntimeDto", string(data)}, " ")
}
