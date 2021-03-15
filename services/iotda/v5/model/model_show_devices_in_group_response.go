package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDevicesInGroupResponse struct {
	// 设备列表。
	Devices        *[]SimplifyDevice `json:"devices,omitempty"`
	Page           *Page             `json:"page,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowDevicesInGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDevicesInGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowDevicesInGroupResponse", string(data)}, " ")
}
