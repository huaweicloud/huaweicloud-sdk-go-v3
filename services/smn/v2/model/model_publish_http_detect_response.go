package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishHttpDetectResponse Response Object
type PublishHttpDetectResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 探测任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishHttpDetectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishHttpDetectResponse struct{}"
	}

	return strings.Join([]string{"PublishHttpDetectResponse", string(data)}, " ")
}
