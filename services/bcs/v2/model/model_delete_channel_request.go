package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteChannelRequest struct {

	// 区块链服务id。可调用“查询服务实例列表”接口获取id
	BlockchainId string `json:"blockchain_id"`

	// 区块链通道名称。可调用“查询实例信息”接口获取，接口返回的“channels”中的name字段值
	ChannelId string `json:"channel_id"`
}

func (o DeleteChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChannelRequest struct{}"
	}

	return strings.Join([]string{"DeleteChannelRequest", string(data)}, " ")
}
