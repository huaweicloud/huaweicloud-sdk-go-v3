package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleTasksResponse Response Object
type ListScheduleTasksResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 定时任务列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]ScheduleTask `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListScheduleTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleTasksResponse", string(data)}, " ")
}
