package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEditTaskResponse Response Object
type CreateEditTaskResponse struct {

	// 编辑任务Id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEditTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateEditTaskResponse", string(data)}, " ")
}
