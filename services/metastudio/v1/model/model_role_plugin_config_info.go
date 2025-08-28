package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RolePluginConfigInfo 角色插件配置信息。
type RolePluginConfigInfo struct {
	PluginType *PluginTypeEnum `json:"plugin_type"`

	PluginSource *PluginSourceEnum `json:"plugin_source"`

	// 插件配置ID。
	PluginConfigId *string `json:"plugin_config_id,omitempty"`
}

func (o RolePluginConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RolePluginConfigInfo struct{}"
	}

	return strings.Join([]string{"RolePluginConfigInfo", string(data)}, " ")
}
