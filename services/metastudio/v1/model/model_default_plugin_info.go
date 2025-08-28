package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefaultPluginInfo 默认插件支持信息。
type DefaultPluginInfo struct {
	PluginType *PluginTypeEnum `json:"plugin_type,omitempty"`

	// 支持默认插件。
	SupportDefault *bool `json:"support_default,omitempty"`
}

func (o DefaultPluginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultPluginInfo struct{}"
	}

	return strings.Join([]string{"DefaultPluginInfo", string(data)}, " ")
}
