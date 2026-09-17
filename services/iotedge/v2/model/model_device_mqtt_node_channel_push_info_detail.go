package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceMqttNodeChannelPushInfoDetail 设备数据推送MQTT详情
type DeviceMqttNodeChannelPushInfoDetail struct {

	// client推送的topic
	Topic *string `json:"topic,omitempty"`

	// Mqtt的服务质量
	Qos *int32 `json:"qos,omitempty"`
}

func (o DeviceMqttNodeChannelPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceMqttNodeChannelPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DeviceMqttNodeChannelPushInfoDetail", string(data)}, " ")
}
