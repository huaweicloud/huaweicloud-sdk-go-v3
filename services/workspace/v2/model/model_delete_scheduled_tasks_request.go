package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduledTasksRequest Request Object
type DeleteScheduledTasksRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`
}

func (o DeleteScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduledTasksRequest", string(data)}, " ")
}
