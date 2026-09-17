package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMqttNodeChannelPushInfoDto MQTT推送信息详情
type UpdateMqttNodeChannelPushInfoDto struct {
	DeviceData *DeviceMqttNodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o UpdateMqttNodeChannelPushInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqttNodeChannelPushInfoDto struct{}"
	}

	return strings.Join([]string{"UpdateMqttNodeChannelPushInfoDto", string(data)}, " ")
}
