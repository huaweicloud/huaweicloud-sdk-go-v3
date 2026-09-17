package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttChannelDetailDto MQTT通道详情
type MqttChannelDetailDto struct {
	ConnectionInfo *MqttConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *MqttPushInfoResp `json:"push_info,omitempty"`
}

func (o MqttChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttChannelDetailDto struct{}"
	}

	return strings.Join([]string{"MqttChannelDetailDto", string(data)}, " ")
}
