package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDeviceGroupsResponse struct {
	// 设备组信息列表。

	DeviceGroups *[]DeviceGroupResponseDto `json:"device_groups,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDeviceGroupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDeviceGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceGroupsResponse", string(data)}, " ")
}
