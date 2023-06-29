package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScheduleTaskResponse Response Object
type CancelScheduleTaskResponse struct {

	// 任务ID，表示成功取消定时调度任务的ID。
	JobIds         *[]string `json:"job_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CancelScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelScheduleTaskResponse", string(data)}, " ")
}
