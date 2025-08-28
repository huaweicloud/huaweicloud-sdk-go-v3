package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginConfigReq 创建插件配置请求。
type CreatePluginConfigReq struct {
	PluginProvider *PluginProviderEnum `json:"plugin_provider"`

	// 密钥。
	ApiKey *string `json:"api_key,omitempty"`
}

func (o CreatePluginConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginConfigReq struct{}"
	}

	return strings.Join([]string{"CreatePluginConfigReq", string(data)}, " ")
}
