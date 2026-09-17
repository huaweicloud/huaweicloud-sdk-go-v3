package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePulsarNodeChannelPushInfoDto Pulsar推送信息详情
type UpdatePulsarNodeChannelPushInfoDto struct {
	DeviceData *DevicePulsarNodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o UpdatePulsarNodeChannelPushInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePulsarNodeChannelPushInfoDto struct{}"
	}

	return strings.Join([]string{"UpdatePulsarNodeChannelPushInfoDto", string(data)}, " ")
}
