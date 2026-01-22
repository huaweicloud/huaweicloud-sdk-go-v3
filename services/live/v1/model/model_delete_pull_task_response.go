package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePullTaskResponse Response Object
type DeletePullTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePullTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePullTaskResponse struct{}"
	}

	return strings.Join([]string{"DeletePullTaskResponse", string(data)}, " ")
}
