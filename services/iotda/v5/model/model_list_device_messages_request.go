package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDeviceMessagesRequest struct {
	DeviceId   string  `json:"device_id"`
	InstanceId *string `json:"Instance-Id,omitempty"`
}

func (o ListDeviceMessagesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDeviceMessagesRequest struct{}"
	}

	return strings.Join([]string{"ListDeviceMessagesRequest", string(data)}, " ")
}
