package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMultiTasksResponse struct {

	// 任务ID, 可为空
	TaskId         *string `json:"task_id,omitempty" xml:"task_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMultiTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateMultiTasksResponse", string(data)}, " ")
}
