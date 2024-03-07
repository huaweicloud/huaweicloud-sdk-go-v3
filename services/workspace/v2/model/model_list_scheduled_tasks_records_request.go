package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRecordsRequest Request Object
type ListScheduledTasksRecordsRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`
}

func (o ListScheduledTasksRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRecordsRequest", string(data)}, " ")
}
