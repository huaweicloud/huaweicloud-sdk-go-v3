package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttNodeChannelPushInfoRsp MQTT推送信息详情
type MqttNodeChannelPushInfoRsp struct {
	DeviceData *DeviceMqttNodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o MqttNodeChannelPushInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttNodeChannelPushInfoRsp struct{}"
	}

	return strings.Join([]string{"MqttNodeChannelPushInfoRsp", string(data)}, " ")
}
