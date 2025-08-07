package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleTasksResponse Response Object
type ListScheduleTasksResponse struct {

	// 定时任务列表。
	ScheduleTasks *[]ScheduleTaskInfo `json:"schedule_tasks,omitempty"`

	// 总任务数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduleTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleTasksResponse", string(data)}, " ")
}
