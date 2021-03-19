package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateOrDeleteDeviceInGroupRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	GroupId string `json:"group_id"`

	ActionId string `json:"action_id"`

	DeviceId string `json:"device_id"`
}

func (o CreateOrDeleteDeviceInGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOrDeleteDeviceInGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateOrDeleteDeviceInGroupRequest", string(data)}, " ")
}
