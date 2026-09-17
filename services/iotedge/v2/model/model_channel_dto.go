package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChannelDto 外部推送路由信息返回结构体
type ChannelDto struct {

	// 规则ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 通道
	Channel *string `json:"channel,omitempty"`

	// 推送地址信息
	Endpoint *string `json:"endpoint,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ChannelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelDto struct{}"
	}

	return strings.Join([]string{"ChannelDto", string(data)}, " ")
}
