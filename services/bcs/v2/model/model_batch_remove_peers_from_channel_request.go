package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemovePeersFromChannelRequest Request Object
type BatchRemovePeersFromChannelRequest struct {

	// 区块链服务id。可调用“查询服务实例列表”接口获取ID
	BlockchainId string `json:"blockchain_id"`

	// 区块链通道名称。可调用“查询实例信息”接口获取，接口返回的“channels”中的name字段值
	ChannelId string `json:"channel_id"`

	Body *BatchRemovePeersFromChannelRequestBody `json:"body,omitempty"`
}

func (o BatchRemovePeersFromChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemovePeersFromChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchRemovePeersFromChannelRequest", string(data)}, " ")
}
