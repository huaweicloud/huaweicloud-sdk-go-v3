package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevicePulsarPushInfoDetail 设备数据推送Pulsar详情
type DevicePulsarPushInfoDetail struct {

	// client推送的topic
	Topic *string `json:"topic,omitempty"`
}

func (o DevicePulsarPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicePulsarPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DevicePulsarPushInfoDetail", string(data)}, " ")
}
