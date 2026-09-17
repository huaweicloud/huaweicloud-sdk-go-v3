package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePulsarNodeChannelDetail 更新Pulsar推送通道详情请求结构体
type UpdatePulsarNodeChannelDetail struct {
	ConnectionInfo *UpdatePulsarNodeChannelConnectionInfo `json:"connection_info,omitempty"`

	PushInfo *UpdatePulsarNodeChannelPushInfoDto `json:"push_info,omitempty"`
}

func (o UpdatePulsarNodeChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePulsarNodeChannelDetail struct{}"
	}

	return strings.Join([]string{"UpdatePulsarNodeChannelDetail", string(data)}, " ")
}
