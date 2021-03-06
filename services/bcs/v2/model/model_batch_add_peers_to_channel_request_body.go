package model

import (
	"encoding/json"

	"strings"
)

// 节点加入通道
type BatchAddPeersToChannelRequestBody struct {
	// 加入某个通道的节点信息

	ChannelPeers []BatchAddPeersToChannelRequestBodyChannelPeers `json:"channel_peers"`
}

func (o BatchAddPeersToChannelRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelRequestBody", string(data)}, " ")
}
