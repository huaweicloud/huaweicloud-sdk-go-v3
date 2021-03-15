package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetDeviceSecretRequest struct {
	InstanceId *string            `json:"Instance-Id,omitempty"`
	DeviceId   string             `json:"device_id"`
	ActionId   string             `json:"action_id"`
	Body       *ResetDeviceSecret `json:"body,omitempty"`
}

func (o ResetDeviceSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetDeviceSecretRequest struct{}"
	}

	return strings.Join([]string{"ResetDeviceSecretRequest", string(data)}, " ")
}
