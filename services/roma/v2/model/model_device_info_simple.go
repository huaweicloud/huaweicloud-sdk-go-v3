package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeviceInfoSimple struct {

	// 设备ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 设备名称
	DeviceName *string `json:"device_name,omitempty" xml:"device_name"`
}

func (o DeviceInfoSimple) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceInfoSimple struct{}"
	}

	return strings.Join([]string{"DeviceInfoSimple", string(data)}, " ")
}
