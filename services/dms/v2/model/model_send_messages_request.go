package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendMessagesRequest struct {
	// 指定的队列ID。

	QueueId string `json:"queue_id"`

	Body *SendMessagesReq `json:"body,omitempty"`
}

func (o SendMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessagesRequest struct{}"
	}

	return strings.Join([]string{"SendMessagesRequest", string(data)}, " ")
}
