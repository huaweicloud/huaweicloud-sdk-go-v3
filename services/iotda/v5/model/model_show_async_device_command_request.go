package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAsyncDeviceCommandRequest struct {
	DeviceId   string  `json:"device_id"`
	InstanceId *string `json:"Instance-Id,omitempty"`
	CommandId  string  `json:"command_id"`
}

func (o ShowAsyncDeviceCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAsyncDeviceCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowAsyncDeviceCommandRequest", string(data)}, " ")
}
