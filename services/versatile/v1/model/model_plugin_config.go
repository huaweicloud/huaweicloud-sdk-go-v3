package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginConfig 插件参数
type PluginConfig struct {

	// 插件id
	PluginId *string `json:"plugin_id,omitempty"`

	// 单个插件参数
	Config map[string]interface{} `json:"config,omitempty"`
}

func (o PluginConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginConfig struct{}"
	}

	return strings.Join([]string{"PluginConfig", string(data)}, " ")
}
