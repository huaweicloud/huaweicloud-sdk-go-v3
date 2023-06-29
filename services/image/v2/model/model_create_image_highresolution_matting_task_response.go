package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageHighresolutionMattingTaskResponse Response Object
type CreateImageHighresolutionMattingTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageHighresolutionMattingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageHighresolutionMattingTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateImageHighresolutionMattingTaskResponse", string(data)}, " ")
}
