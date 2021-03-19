package model

import (
	"encoding/json"

	"strings"
)

type ResetDeviceSecret struct {
	// 设备密钥，设置该字段时平台将设备密钥重置为指定值，若不设置则由平台自动生成。

	Secret *string `json:"secret,omitempty"`
	// 是否强制断开设备的连接，当前仅限长连接。默认值false。

	ForceDisconnect *bool `json:"force_disconnect,omitempty"`
}

func (o ResetDeviceSecret) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetDeviceSecret struct{}"
	}

	return strings.Join([]string{"ResetDeviceSecret", string(data)}, " ")
}
