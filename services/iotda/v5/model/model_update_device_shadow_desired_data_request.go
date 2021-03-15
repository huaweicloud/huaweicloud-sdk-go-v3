package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDeviceShadowDesiredDataRequest struct {
	InstanceId *string         `json:"Instance-Id,omitempty"`
	DeviceId   string          `json:"device_id"`
	Body       *UpdateDesireds `json:"body,omitempty"`
}

func (o UpdateDeviceShadowDesiredDataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDeviceShadowDesiredDataRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceShadowDesiredDataRequest", string(data)}, " ")
}
