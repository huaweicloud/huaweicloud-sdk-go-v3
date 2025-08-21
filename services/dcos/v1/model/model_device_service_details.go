package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceServiceDetails 操作设备对象的服务详情
type DeviceServiceDetails struct {

	// 对象类型
	Type *string `json:"type,omitempty"`

	// 机房编码
	RoomCode *string `json:"room_code,omitempty"`

	// 机柜编码
	RackCode *string `json:"rack_code,omitempty"`

	// u位
	UPosition *string `json:"u_position,omitempty"`

	// sn
	Sn *string `json:"sn,omitempty"`

	// 端口号或槽口号
	PortOrSlot *string `json:"port_or_slot,omitempty"`

	// 成功或异常
	TaskStatus *string `json:"task_status,omitempty"`

	// 异常原因
	Description *string `json:"description,omitempty"`
}

func (o DeviceServiceDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceServiceDetails struct{}"
	}

	return strings.Join([]string{"DeviceServiceDetails", string(data)}, " ")
}
