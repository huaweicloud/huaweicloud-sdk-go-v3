package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	DeviceId string `json:"device_id"`

	Body *UpdateDevice `json:"body,omitempty"`
}

func (o UpdateDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequest", string(data)}, " ")
}
