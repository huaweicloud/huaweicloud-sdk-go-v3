package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsTaskResponse Response Object
type CreateFsTaskResponse struct {

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateFsTaskResponse", string(data)}, " ")
}
