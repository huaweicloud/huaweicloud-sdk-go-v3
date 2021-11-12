package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建通道的配置
type BatchCreateChannelsRequestBody struct {
	// 通道列表

	Channels []ChannelCreateInfo `json:"channels"`
}

func (o BatchCreateChannelsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsRequestBody", string(data)}, " ")
}
