package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDeviceGroupRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	GroupId string `json:"group_id"`
}

func (o ShowDeviceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupRequest", string(data)}, " ")
}
