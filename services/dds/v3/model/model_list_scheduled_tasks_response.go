package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksResponse Response Object
type ListScheduledTasksResponse struct {

	// 记录数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 任务详情
	Schedules      *[]ScheduledTasksRspSchedules `json:"schedules,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListScheduledTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksResponse", string(data)}, " ")
}
