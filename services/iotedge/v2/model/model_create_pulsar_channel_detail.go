package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePulsarChannelDetail Pulsar通道配置详情
type CreatePulsarChannelDetail struct {
	ConnectionInfo *PulsarConnectionInfo `json:"connection_info"`

	PushInfo *PulsarPushInfo `json:"push_info"`
}

func (o CreatePulsarChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePulsarChannelDetail struct{}"
	}

	return strings.Join([]string{"CreatePulsarChannelDetail", string(data)}, " ")
}
