package model

import (
	"encoding/json"

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
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmConsumptionMessagesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmConsumptionMessagesRequest", string(data)}, " ")
}
