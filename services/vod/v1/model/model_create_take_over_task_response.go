package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTakeOverTaskResponse struct {
	// 任务ID。

	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTakeOverTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTakeOverTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTakeOverTaskResponse", string(data)}, " ")
}
