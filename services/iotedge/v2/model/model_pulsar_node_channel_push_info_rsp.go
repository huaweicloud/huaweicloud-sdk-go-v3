package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarNodeChannelPushInfoRsp Pulsar推送信息详情
type PulsarNodeChannelPushInfoRsp struct {
	DeviceData *DevicePulsarNodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o PulsarNodeChannelPushInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarNodeChannelPushInfoRsp struct{}"
	}

	return strings.Join([]string{"PulsarNodeChannelPushInfoRsp", string(data)}, " ")
}
