package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddDeviceGroupRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddDeviceGroupDto `json:"body,omitempty"`
}

func (o AddDeviceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"AddDeviceGroupRequest", string(data)}, " ")
}
