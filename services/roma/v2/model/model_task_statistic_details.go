package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 每种状态任务的数量信息
type TaskStatisticDetails struct {

	// 运行任务数量
	RunningTasksCount *int32 `json:"running_tasks_count,omitempty" xml:"running_tasks_count"`

	// 异常任务数量
	AbnormalTasksCount *int32 `json:"abnormal_tasks_count,omitempty" xml:"abnormal_tasks_count"`

	// 终止任务数量
	TerminatedTasksCount *int32 `json:"terminated_tasks_count,omitempty" xml:"terminated_tasks_count"`

	// 成功任务数量
	CompletedTasksCount *int32 `json:"completed_tasks_count,omitempty" xml:"completed_tasks_count"`
}

func (o TaskStatisticDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatisticDetails struct{}"
	}

	return strings.Join([]string{"TaskStatisticDetails", string(data)}, " ")
}
