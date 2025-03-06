package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleTaskResponse Response Object
type ListScheduleTaskResponse struct {

	// 任务列表。
	Tasks *[]ScheduleTaskDetail `json:"tasks,omitempty"`

	// 任务数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleTaskResponse", string(data)}, " ")
}
