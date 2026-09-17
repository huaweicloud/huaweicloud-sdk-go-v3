package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarPushInfo 创建Pulsar推送信息
type PulsarPushInfo struct {
	DeviceData *DevicePulsarPushInfo `json:"device_data,omitempty"`
}

func (o PulsarPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarPushInfo struct{}"
	}

	return strings.Join([]string{"PulsarPushInfo", string(data)}, " ")
}
