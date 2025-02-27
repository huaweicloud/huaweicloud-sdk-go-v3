package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleRecordTasksResponse Response Object
type CreateScheduleRecordTasksResponse struct {

	// 创建成功的任务ID
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScheduleRecordTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleRecordTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduleRecordTasksResponse", string(data)}, " ")
}
