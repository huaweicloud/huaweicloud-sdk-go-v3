package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteQueueRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	QueueId    string  `json:"queue_id"`
}

func (o DeleteQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueRequest", string(data)}, " ")
}
