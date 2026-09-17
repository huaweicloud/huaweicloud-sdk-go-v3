package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevicePulsarNodeChannelPushInfoDetail 设备数据推送Pulsar详情
type DevicePulsarNodeChannelPushInfoDetail struct {

	// client推送的topic
	Topic *string `json:"topic,omitempty"`
}

func (o DevicePulsarNodeChannelPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicePulsarNodeChannelPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DevicePulsarNodeChannelPushInfoDetail", string(data)}, " ")
}
