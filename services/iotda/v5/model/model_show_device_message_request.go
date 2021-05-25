package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDeviceMessageRequest struct {
	// 下发消息的设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获。

	DeviceId string `json:"device_id"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 下发消息的消息ID，用于唯一标识一个消息，在消息下发时由物联网平台分配获得。

	MessageId string `json:"message_id"`
}

func (o ShowDeviceMessageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDeviceMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceMessageRequest", string(data)}, " ")
}
