package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDeviceGroupRequest struct {
	InstanceId *string               `json:"Instance-Id,omitempty"`
	GroupId    string                `json:"group_id"`
	Body       *UpdateDeviceGroupDto `json:"body,omitempty"`
}

func (o UpdateDeviceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupRequest", string(data)}, " ")
}
