package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMqttNodeChannelDetail 更新MQTT推送通道详情请求结构体
type UpdateMqttNodeChannelDetail struct {
	ConnectionInfo *UpdateMqttNodeChannelConnectionInfo `json:"connection_info,omitempty"`

	PushInfo *UpdateMqttNodeChannelPushInfoDto `json:"push_info,omitempty"`
}

func (o UpdateMqttNodeChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqttNodeChannelDetail struct{}"
	}

	return strings.Join([]string{"UpdateMqttNodeChannelDetail", string(data)}, " ")
}
