package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDeviceMessageRequest struct {
	DeviceId   string  `json:"device_id"`
	InstanceId *string `json:"Instance-Id,omitempty"`
	MessageId  string  `json:"message_id"`
}

func (o ShowDeviceMessageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDeviceMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceMessageRequest", string(data)}, " ")
}
