package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateConsumerGroupRequest struct {
	// 指定的队列ID

	QueueId string `json:"queue_id"`

	Body *CreateConsumerGroupReq `json:"body,omitempty"`
}

func (o CreateConsumerGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupRequest", string(data)}, " ")
}
