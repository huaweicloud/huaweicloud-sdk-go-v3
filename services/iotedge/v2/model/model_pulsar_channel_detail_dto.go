package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarChannelDetailDto Pulsar通道详情
type PulsarChannelDetailDto struct {
	ConnectionInfo *PulsarConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *PulsarPushInfoResp `json:"push_info,omitempty"`
}

func (o PulsarChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarChannelDetailDto struct{}"
	}

	return strings.Join([]string{"PulsarChannelDetailDto", string(data)}, " ")
}
