package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddPeersToChannelRequestBody 节点加入通道
type BatchAddPeersToChannelRequestBody struct {

	// 加入某个通道的节点信息
	ChannelPeers []PeerChannelInfo `json:"channel_peers"`
}

func (o BatchAddPeersToChannelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelRequestBody", string(data)}, " ")
}
