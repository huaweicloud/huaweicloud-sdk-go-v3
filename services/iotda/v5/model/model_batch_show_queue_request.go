package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchShowQueueRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	QueueName  *string `json:"queue_name,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	Marker     *string `json:"marker,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
}

func (o BatchShowQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchShowQueueRequest struct{}"
	}

	return strings.Join([]string{"BatchShowQueueRequest", string(data)}, " ")
}
