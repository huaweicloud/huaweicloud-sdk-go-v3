package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchRemoveOrgsFromChannelRequest struct {
	// 区块链服务id。

	BlockchainId string `json:"blockchain_id"`
	// 区块链通道id。

	ChannelId string `json:"channel_id"`

	Body *BatchRemoveOrgsFromChannelRequestBody `json:"body,omitempty"`
}

func (o BatchRemoveOrgsFromChannelRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchRemoveOrgsFromChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveOrgsFromChannelRequest", string(data)}, " ")
}
