package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksResponse Response Object
type ListScheduledTasksResponse struct {

	// 定时任务列表。
	ScheduledTasks *[]ScheduledTask `json:"scheduled_tasks,omitempty"`

	// 总个数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksResponse", string(data)}, " ")
}
