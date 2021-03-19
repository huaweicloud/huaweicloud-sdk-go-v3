package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDeviceGroupRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	GroupId string `json:"group_id"`
}

func (o DeleteDeviceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupRequest", string(data)}, " ")
}
