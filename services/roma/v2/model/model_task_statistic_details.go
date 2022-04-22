package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 每种状态任务的数量信息
type TaskStatisticDetails struct {

	// 运行任务数量
	RunningTasksCount *int32 `json:"running_tasks_count,omitempty"`

	// 异常任务数量
	AbnormalTasksCount *int32 `json:"abnormal_tasks_count,omitempty"`

	// 终止任务数量
	TerminatedTasksCount *int32 `json:"terminated_tasks_count,omitempty"`

	// 成功任务数量
	CompletedTasksCount *int32 `json:"completed_tasks_count,omitempty"`
}

func (o TaskStatisticDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatisticDetails struct{}"
	}

	return strings.Join([]string{"TaskStatisticDetails", string(data)}, " ")
}
