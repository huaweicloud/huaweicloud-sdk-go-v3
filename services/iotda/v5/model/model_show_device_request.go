package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	DeviceId string `json:"device_id"`
}

func (o ShowDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceRequest", string(data)}, " ")
}
