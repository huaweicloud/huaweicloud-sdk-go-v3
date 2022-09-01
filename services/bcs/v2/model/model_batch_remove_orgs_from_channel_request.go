package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchRemoveOrgsFromChannelRequest struct {

	// 区块链服务id。
	BlockchainId string `json:"blockchain_id" xml:"blockchain_id"`

	// 区块链通道名称。
	ChannelId string `json:"channel_id" xml:"channel_id"`

	Body *BatchRemoveOrgsFromChannelRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchRemoveOrgsFromChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveOrgsFromChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveOrgsFromChannelRequest", string(data)}, " ")
}
