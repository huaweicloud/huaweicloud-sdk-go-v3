package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMqttChannelDetail MQTT通道配置详情
type CreateMqttChannelDetail struct {
	ConnectionInfo *ItMqttConnectionInfo `json:"connection_info"`

	PushInfo *MqttPushInfo `json:"push_info"`
}

func (o CreateMqttChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMqttChannelDetail struct{}"
	}

	return strings.Join([]string{"CreateMqttChannelDetail", string(data)}, " ")
}
