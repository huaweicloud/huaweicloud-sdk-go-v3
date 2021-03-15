package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMessageRequest struct {
	DeviceId   string                `json:"device_id"`
	InstanceId *string               `json:"Instance-Id,omitempty"`
	Body       *DeviceMessageRequest `json:"body,omitempty"`
}

func (o CreateMessageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMessageRequest struct{}"
	}

	return strings.Join([]string{"CreateMessageRequest", string(data)}, " ")
}
