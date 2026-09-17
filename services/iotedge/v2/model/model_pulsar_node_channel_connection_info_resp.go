package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarNodeChannelConnectionInfoResp 外部推送通道返回详情
type PulsarNodeChannelConnectionInfoResp struct {

	// 鉴权token
	Token *string `json:"token,omitempty"`
}

func (o PulsarNodeChannelConnectionInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarNodeChannelConnectionInfoResp struct{}"
	}

	return strings.Join([]string{"PulsarNodeChannelConnectionInfoResp", string(data)}, " ")
}
