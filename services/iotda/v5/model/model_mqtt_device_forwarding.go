package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttDeviceForwarding 消息重发布消息结构
type MqttDeviceForwarding struct {

	// **参数说明**：消息重发布topic
	Topic string `json:"topic"`

	// **参数说明**：消息重发布过期时间，单位为分钟
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o MqttDeviceForwarding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttDeviceForwarding struct{}"
	}

	return strings.Join([]string{"MqttDeviceForwarding", string(data)}, " ")
}
