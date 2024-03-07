package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScheduledTasksRequest Request Object
type ShowScheduledTasksRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`
}

func (o ShowScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"ShowScheduledTasksRequest", string(data)}, " ")
}
