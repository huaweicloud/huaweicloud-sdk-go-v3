package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ConfirmConsumptionMessagesRequest struct {

	// 队列ID。
	QueueId string `json:"queue_id"`

	// 消费组ID。
	ConsumerGroupId string `json:"consumer_group_id"`

	Body *ConfirmConsumptionMessagesReq `json:"body,omitempty"`
}

func (o ConfirmConsumptionMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmConsumptionMessagesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmConsumptionMessagesRequest", string(data)}, " ")
}
