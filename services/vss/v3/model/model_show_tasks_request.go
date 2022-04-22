package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTasksRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksRequest struct{}"
	}

	return strings.Join([]string{"ShowTasksRequest", string(data)}, " ")
}
