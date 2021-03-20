package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQueueRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	QueueId string `json:"queue_id"`
}

func (o ShowQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQueueRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueRequest", string(data)}, " ")
}
