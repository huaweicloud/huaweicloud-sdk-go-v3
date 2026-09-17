package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttPushInfo 创建MQTT推送信息
type MqttPushInfo struct {
	DeviceData *DeviceMqttPushInfo `json:"device_data,omitempty"`
}

func (o MqttPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttPushInfo struct{}"
	}

	return strings.Join([]string{"MqttPushInfo", string(data)}, " ")
}
