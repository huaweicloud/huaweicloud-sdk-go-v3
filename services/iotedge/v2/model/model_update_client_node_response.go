package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientNodeResponse Response Object
type UpdateClientNodeResponse struct {

	// 通道
	Channel *string `json:"channel,omitempty"`

	// 推送通道描述
	Description *string `json:"description,omitempty"`

	// 推送的地址
	Endpoint *string `json:"endpoint,omitempty"`

	MqttChannelDetail *MqttNodeChannelDetailDto `json:"mqtt_channel_detail,omitempty"`

	IotdbChannelDetail *IoTdbNodeChannelDetailDto `json:"iotdb_channel_detail,omitempty"`

	Influxdb2ChannelDetail *InfluxDb2NodeChannelDetailDto `json:"influxdb2_channel_detail,omitempty"`

	PulsarChannelDetail *PulsarNodeChannelDetailDto `json:"pulsar_channel_detail,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 下发时间，表示通道是否已经同步到了节点
	SynchronizedTime *string `json:"synchronized_time,omitempty"`

	// 下发状态，表示是否已同步到了节点
	SynchronizedStatus *bool `json:"synchronized_status,omitempty"`
	HttpStatusCode     int   `json:"-"`
}

func (o UpdateClientNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientNodeResponse", string(data)}, " ")
}
