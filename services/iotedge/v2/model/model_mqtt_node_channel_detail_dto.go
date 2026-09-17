package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttNodeChannelDetailDto MQTT通道详情
type MqttNodeChannelDetailDto struct {
	ConnectionInfo *MqttNodeChannelConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *MqttNodeChannelPushInfoRsp `json:"push_info,omitempty"`
}

func (o MqttNodeChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttNodeChannelDetailDto struct{}"
	}

	return strings.Join([]string{"MqttNodeChannelDetailDto", string(data)}, " ")
}
