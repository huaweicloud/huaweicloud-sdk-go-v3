package model

import (
	"encoding/json"

	"strings"
)

// 下发设备命令消息结构
type ActionDeviceCommand struct {
	// 下发命令的设备ID。 - 当创建设备数据规则时，若device_id为空，则命令下发给触发条件的设备。 - 当创建定时规则时，不允许为空。

	DeviceId *string `json:"device_id,omitempty"`

	Cmd *Cmd `json:"cmd"`
}

func (o ActionDeviceCommand) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionDeviceCommand struct{}"
	}

	return strings.Join([]string{"ActionDeviceCommand", string(data)}, " ")
}
