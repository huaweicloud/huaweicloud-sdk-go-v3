package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPullTaskResponse Response Object
type ModifyPullTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyPullTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPullTaskResponse struct{}"
	}

	return strings.Join([]string{"ModifyPullTaskResponse", string(data)}, " ")
}
