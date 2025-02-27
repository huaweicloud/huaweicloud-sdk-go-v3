package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleRecordTasksResponse Response Object
type ListScheduleRecordTasksResponse struct {

	// 查询结果的总元素数量
	Total *int32 `json:"total,omitempty"`

	// 记录偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 当前页条目数
	Limit *int32 `json:"limit,omitempty"`

	// 录制模板数组
	Tasks *[]ScheduleRecordTasks `json:"tasks,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListScheduleRecordTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleRecordTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleRecordTasksResponse", string(data)}, " ")
}
