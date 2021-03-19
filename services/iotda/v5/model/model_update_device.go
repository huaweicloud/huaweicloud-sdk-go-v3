package model

import (
	"encoding/json"

	"strings"
)

// 修改设备信息对象。
type UpdateDevice struct {
	// 设备名称。

	DeviceName *string `json:"device_name,omitempty"`
	// 设备的描述信息。

	Description *string `json:"description,omitempty"`
	// 设备扩展信息。用户可以自定义任何想要的扩展信息，修改子设备信息时不会下发给网关。

	ExtensionInfo *interface{} `json:"extension_info,omitempty"`

	AuthInfo *AuthInfoWithoutSecret `json:"auth_info,omitempty"`
}

func (o UpdateDevice) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDevice struct{}"
	}

	return strings.Join([]string{"UpdateDevice", string(data)}, " ")
}
