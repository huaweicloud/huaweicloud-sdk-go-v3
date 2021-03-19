package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type FreezeDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	DeviceId string `json:"device_id"`
}

func (o FreezeDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FreezeDeviceRequest struct{}"
	}

	return strings.Join([]string{"FreezeDeviceRequest", string(data)}, " ")
}
