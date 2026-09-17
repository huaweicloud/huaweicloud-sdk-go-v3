package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChannelRequestDto 创建外部推送通道请求结构体
type CreateChannelRequestDto struct {

	// 推送通道ID,非必填，若用户不填，则系统自动生成
	ChannelId *string `json:"channel_id,omitempty"`

	// 通道
	Channel string `json:"channel"`

	// 推送通道名称
	Name string `json:"name"`

	// 推送通道描述
	Description *string `json:"description,omitempty"`

	// 推送的地址
	Endpoint string `json:"endpoint"`

	MqttChannelDetail *CreateMqttChannelDetail `json:"mqtt_channel_detail,omitempty"`

	IotdbChannelDetail *CreateIoTdbChannelDetail `json:"iotdb_channel_detail,omitempty"`

	Influxdb2ChannelDetail *CreateInfluxDb2ChannelDetail `json:"influxdb2_channel_detail,omitempty"`

	PulsarChannelDetail *CreatePulsarChannelDetail `json:"pulsar_channel_detail,omitempty"`
}

func (o CreateChannelRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChannelRequestDto struct{}"
	}

	return strings.Join([]string{"CreateChannelRequestDto", string(data)}, " ")
}
