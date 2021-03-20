package model

import (
	"encoding/json"

	"strings"
)

// 条件中设备状态类型的信息，自定义结构。
type DeviceStatusCondition struct {
	// 状态列表，设备状态条件携带该参数。

	StatusList *[]string `json:"status_list,omitempty"`
}

func (o DeviceStatusCondition) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeviceStatusCondition struct{}"
	}

	return strings.Join([]string{"DeviceStatusCondition", string(data)}, " ")
}
