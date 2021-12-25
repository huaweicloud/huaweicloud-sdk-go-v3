package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevicesInGroup struct {
	// 设备ID

	DeviceId *int32 `json:"device_id,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 设备状态 0-启用 1-禁用

	Status *int32 `json:"status,omitempty"`
	// 是否在线 0-未连接 1-在线 2-离线

	OnlineStatus *int32 `json:"online_status,omitempty"`
}

func (o DevicesInGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicesInGroup struct{}"
	}

	return strings.Join([]string{"DevicesInGroup", string(data)}, " ")
}
