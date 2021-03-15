package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UnfreezeDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	DeviceId   string  `json:"device_id"`
}

func (o UnfreezeDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnfreezeDeviceRequest struct{}"
	}

	return strings.Join([]string{"UnfreezeDeviceRequest", string(data)}, " ")
}
