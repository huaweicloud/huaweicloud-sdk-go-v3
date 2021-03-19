package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDevicesInGroupRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	GroupId string `json:"group_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowDevicesInGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDevicesInGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDevicesInGroupRequest", string(data)}, " ")
}
