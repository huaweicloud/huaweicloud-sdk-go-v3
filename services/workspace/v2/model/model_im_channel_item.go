package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImChannelItem IM 通道配置项（不含密钥）
type ImChannelItem struct {

	// IM 平台类型：wecom / feishu / dingtalk-connector
	Platform string `json:"platform"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 客户端 ID
	ClientId *string `json:"client_id,omitempty"`

	// 平台扩展配置
	PlatformSpecific map[string]string `json:"platform_specific,omitempty"`
}

func (o ImChannelItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImChannelItem struct{}"
	}

	return strings.Join([]string{"ImChannelItem", string(data)}, " ")
}
