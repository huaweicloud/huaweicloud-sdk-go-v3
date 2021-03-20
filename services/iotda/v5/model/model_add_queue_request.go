package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddQueueRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *QueueInfo `json:"body,omitempty"`
}

func (o AddQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddQueueRequest struct{}"
	}

	return strings.Join([]string{"AddQueueRequest", string(data)}, " ")
}
