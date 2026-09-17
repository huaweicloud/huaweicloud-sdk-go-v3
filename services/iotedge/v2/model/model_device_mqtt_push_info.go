package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceMqttPushInfo 创建设备数据推送MQTT
type DeviceMqttPushInfo struct {

	// client推送的topic
	Topic string `json:"topic"`

	// 数据格式转换类型
	Format *string `json:"format,omitempty"`

	// Mqtt的服务质量
	Qos *int32 `json:"qos,omitempty"`
}

func (o DeviceMqttPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceMqttPushInfo struct{}"
	}

	return strings.Join([]string{"DeviceMqttPushInfo", string(data)}, " ")
}
