package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeerChannelInfo struct {
	// peer加入的通道名称

	ChannelName string `json:"channel_name"`
	// 加入通道peer名称和数量，key为组织名称，value为peer数量

	Peers map[string]int64 `json:"peers"`
}

func (o PeerChannelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerChannelInfo struct{}"
	}

	return strings.Join([]string{"PeerChannelInfo", string(data)}, " ")
}
