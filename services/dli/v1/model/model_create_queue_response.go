package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueueResponse Response Object
type CreateQueueResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 新增队列的名称。
	QueueName      *string `json:"queue_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueResponse struct{}"
	}

	return strings.Join([]string{"CreateQueueResponse", string(data)}, " ")
}
