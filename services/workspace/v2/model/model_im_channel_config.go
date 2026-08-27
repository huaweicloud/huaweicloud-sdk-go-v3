package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImChannelConfig IM 通道配置
type ImChannelConfig struct {

	// IM 平台类型：wecom / feishu / dingtalk-connector
	Platform string `json:"platform"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 客户端 ID
	ClientId *string `json:"client_id,omitempty"`

	// 客户端密钥
	ClientSecret *string `json:"client_secret,omitempty"`

	// 平台扩展配置
	PlatformSpecific map[string]string `json:"platform_specific,omitempty"`
}

func (o ImChannelConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImChannelConfig struct{}"
	}

	return strings.Join([]string{"ImChannelConfig", string(data)}, " ")
}
