package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeChannelRequestDto 更新已分配节点通道请求结构体
type UpdateNodeChannelRequestDto struct {
	MqttChannelDetail *UpdateMqttNodeChannelDetail `json:"mqtt_channel_detail,omitempty"`

	PulsarChannelDetail *UpdatePulsarNodeChannelDetail `json:"pulsar_channel_detail,omitempty"`
}

func (o UpdateNodeChannelRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeChannelRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateNodeChannelRequestDto", string(data)}, " ")
}
