package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttPushInfoResp MQTT推送信息详情
type MqttPushInfoResp struct {
	DeviceData *DeviceMqttPushInfoDetail `json:"device_data,omitempty"`
}

func (o MqttPushInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttPushInfoResp struct{}"
	}

	return strings.Join([]string{"MqttPushInfoResp", string(data)}, " ")
}
