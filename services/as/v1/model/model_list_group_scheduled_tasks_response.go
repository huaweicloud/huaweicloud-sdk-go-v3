package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupScheduledTasksResponse Response Object
type ListGroupScheduledTasksResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 计划任务列表
	ScheduledTasks *[]ScheduledTaskDetail `json:"scheduled_tasks,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListGroupScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"ListGroupScheduledTasksResponse", string(data)}, " ")
}
