package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDevicesResponse struct {
	// 设备信息列表。

	Devices *[]QueryDeviceSimplify `json:"devices,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
