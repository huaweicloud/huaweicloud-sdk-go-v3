package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarNodeChannelDetailDto Pulsar通道详情
type PulsarNodeChannelDetailDto struct {
	ConnectionInfo *PulsarNodeChannelConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *PulsarNodeChannelPushInfoRsp `json:"push_info,omitempty"`
}

func (o PulsarNodeChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarNodeChannelDetailDto struct{}"
	}

	return strings.Join([]string{"PulsarNodeChannelDetailDto", string(data)}, " ")
}
