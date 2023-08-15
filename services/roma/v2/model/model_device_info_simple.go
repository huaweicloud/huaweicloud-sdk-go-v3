package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeviceInfoSimple struct {

	// 设备ID
	Id *int32 `json:"id,omitempty"`

	// 设备名称
	DeviceName *string `json:"device_name,omitempty"`

	// 设备ID（兼容20.0）
	DeviceId *int64 `json:"device_id,omitempty"`
}

func (o DeviceInfoSimple) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceInfoSimple struct{}"
	}

	return strings.Join([]string{"DeviceInfoSimple", string(data)}, " ")
}
