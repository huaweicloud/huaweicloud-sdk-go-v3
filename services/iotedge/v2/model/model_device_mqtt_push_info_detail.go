package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceMqttPushInfoDetail 设备数据推送MQTT详情
type DeviceMqttPushInfoDetail struct {

	// client推送的topic
	Topic *string `json:"topic,omitempty"`

	// 数据格式转换类型
	Format *string `json:"format,omitempty"`

	// MQTT的服务质量
	Qos *int32 `json:"qos,omitempty"`
}

func (o DeviceMqttPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceMqttPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DeviceMqttPushInfoDetail", string(data)}, " ")
}
