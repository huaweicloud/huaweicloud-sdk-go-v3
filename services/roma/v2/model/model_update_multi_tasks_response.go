package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMultiTasksResponse struct {

	// 任务ID, 可为空
	TaskId         *string `json:"task_id,omitempty" xml:"task_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMultiTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiTasksResponse struct{}"
	}

	return strings.Join([]string{"UpdateMultiTasksResponse", string(data)}, " ")
}
