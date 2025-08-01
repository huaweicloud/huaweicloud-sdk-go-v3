package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskResponse Response Object
type CreateTaskResponse struct {

	// 本次创建任务的id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTaskResponse", string(data)}, " ")
}
