package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MQTT配置。
type EdgeNodeMqttConfig struct {

	// 是否启用MQTT
	EnableMqtt bool `json:"enable_mqtt"`

	// MQTT配置 当enable_mqtt取值为false时，mqtts需要为空数组
	Mqtts []Mqtt `json:"mqtts"`
}

func (o EdgeNodeMqttConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeMqttConfig struct{}"
	}

	return strings.Join([]string{"EdgeNodeMqttConfig", string(data)}, " ")
}
