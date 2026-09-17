package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePushChannelResponse Response Object
type CreatePushChannelResponse struct {

	// 推送通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 推送通道名称
	Name *string `json:"name,omitempty"`

	// 通道
	Channel *string `json:"channel,omitempty"`

	// 推送通道描述
	Description *string `json:"description,omitempty"`

	// 推送的地址
	Endpoint *string `json:"endpoint,omitempty"`

	MqttChannelDetail *MqttChannelDetailDto `json:"mqtt_channel_detail,omitempty"`

	IotdbChannelDetail *IoTdbChannelDetailDto `json:"iotdb_channel_detail,omitempty"`

	Influxdb2ChannelDetail *CreateInfluxDb2ChannelDetail `json:"influxdb2_channel_detail,omitempty"`

	PulsarChannelDetail *PulsarChannelDetailDto `json:"pulsar_channel_detail,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePushChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePushChannelResponse struct{}"
	}

	return strings.Join([]string{"CreatePushChannelResponse", string(data)}, " ")
}
