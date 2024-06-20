package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryDto struct {

	// 插件名
	PluginName string `json:"plugin_name"`

	// 展示名
	DisplayName string `json:"display_name"`

	// 版本号
	Version string `json:"version"`

	// 插件属性,custom/official
	PluginAttribution string `json:"plugin_attribution"`

	// 版本属性,draft/formal
	VersionAttribution *string `json:"version_attribution,omitempty"`
}

func (o PluginPartQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryDto struct{}"
	}

	return strings.Join([]string{"PluginPartQueryDto", string(data)}, " ")
}
