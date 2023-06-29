package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemTasksRequest Request Object
type ListSystemTasksRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 系统任务id.
	TaskId string `json:"task_id"`
}

func (o ListSystemTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSystemTasksRequest", string(data)}, " ")
}
