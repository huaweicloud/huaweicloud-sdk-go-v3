package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTextToImageTaskResponse struct {

	// 任务唯一标识
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTextToImageTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTextToImageTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTextToImageTaskResponse", string(data)}, " ")
}
