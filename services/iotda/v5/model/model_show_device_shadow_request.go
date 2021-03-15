package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDeviceShadowRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	DeviceId   string  `json:"device_id"`
}

func (o ShowDeviceShadowRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDeviceShadowRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceShadowRequest", string(data)}, " ")
}
