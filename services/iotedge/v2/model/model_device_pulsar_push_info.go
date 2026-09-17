package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevicePulsarPushInfo 创建设备数据推送Pulsar
type DevicePulsarPushInfo struct {

	// client推送的topic
	Topic string `json:"topic"`
}

func (o DevicePulsarPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicePulsarPushInfo struct{}"
	}

	return strings.Join([]string{"DevicePulsarPushInfo", string(data)}, " ")
}
