package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginConfigInfo 插件配置基本信息。
type PluginConfigInfo struct {

	// 插件配置ID。
	PluginConfigId *string `json:"plugin_config_id,omitempty"`

	PluginProvider *PluginProviderEnum `json:"plugin_provider,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o PluginConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginConfigInfo struct{}"
	}

	return strings.Join([]string{"PluginConfigInfo", string(data)}, " ")
}
