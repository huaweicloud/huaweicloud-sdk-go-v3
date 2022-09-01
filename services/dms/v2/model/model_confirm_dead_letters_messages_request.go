package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ConfirmDeadLettersMessagesRequest struct {

	// 队列ID。
	QueueId string `json:"queue_id" xml:"queue_id"`

	// 消费组ID。
	ConsumerGroupId string `json:"consumer_group_id" xml:"consumer_group_id"`

	Body *ConfirmDeadLettersMessagesReq `json:"body,omitempty" xml:"body"`
}

func (o ConfirmDeadLettersMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDeadLettersMessagesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmDeadLettersMessagesRequest", string(data)}, " ")
}
