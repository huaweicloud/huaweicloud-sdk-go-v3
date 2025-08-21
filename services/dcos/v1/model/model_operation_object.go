package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperationObject 操作对象
type OperationObject struct {

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
}

func (o OperationObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationObject struct{}"
	}

	return strings.Join([]string{"OperationObject", string(data)}, " ")
}
