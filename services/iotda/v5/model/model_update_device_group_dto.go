package model

import (
	"encoding/json"

	"strings"
)

// 更新设备组请求结构体
type UpdateDeviceGroupDto struct {
	// 设备组名称，单个资源空间下不可重复。

	Name *string `json:"name,omitempty"`
	// 设备组描述。

	Description *string `json:"description,omitempty"`
}

func (o UpdateDeviceGroupDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupDto struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupDto", string(data)}, " ")
}
