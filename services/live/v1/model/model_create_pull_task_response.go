package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePullTaskResponse Response Object
type CreatePullTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePullTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePullTaskResponse struct{}"
	}

	return strings.Join([]string{"CreatePullTaskResponse", string(data)}, " ")
}
