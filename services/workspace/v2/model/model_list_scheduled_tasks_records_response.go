package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRecordsResponse Response Object
type ListScheduledTasksRecordsResponse struct {

	// 定时任务执行记录列表。
	TasksRecords *[]ScheduledTasksRecords `json:"tasks_records,omitempty"`

	// 总个数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduledTasksRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRecordsResponse", string(data)}, " ")
}
