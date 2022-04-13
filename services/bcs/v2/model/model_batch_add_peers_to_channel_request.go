package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddPeersToChannelRequest struct {
	// bcs 服务id

	BlockchainId string `json:"blockchain_id"`

	Body *BatchAddPeersToChannelRequestBody `json:"body,omitempty"`
}

func (o BatchAddPeersToChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelRequest", string(data)}, " ")
}
