package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveImChannelsReq IM 通道配置请求
type SaveImChannelsReq struct {

	// Agent 实例主键 ID
	Id string `json:"id"`

	// IM 通道配置列表
	ImChannels []ImChannelConfig `json:"im_channels"`
}

func (o SaveImChannelsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveImChannelsReq struct{}"
	}

	return strings.Join([]string{"SaveImChannelsReq", string(data)}, " ")
}
