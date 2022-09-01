package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSystemTasksRequest struct {

	// 系统任务id.
	TaskId string `json:"task_id" xml:"task_id"`
}

func (o ListSystemTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSystemTasksRequest", string(data)}, " ")
}
