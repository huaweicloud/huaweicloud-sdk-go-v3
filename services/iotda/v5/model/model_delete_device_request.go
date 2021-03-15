package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	DeviceId   string  `json:"device_id"`
}

func (o DeleteDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceRequest", string(data)}, " ")
}
