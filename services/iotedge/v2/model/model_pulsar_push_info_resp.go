package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarPushInfoResp Pulsar推送信息详情
type PulsarPushInfoResp struct {
	DeviceData *DevicePulsarPushInfoDetail `json:"device_data,omitempty"`
}

func (o PulsarPushInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarPushInfoResp struct{}"
	}

	return strings.Join([]string{"PulsarPushInfoResp", string(data)}, " ")
}
