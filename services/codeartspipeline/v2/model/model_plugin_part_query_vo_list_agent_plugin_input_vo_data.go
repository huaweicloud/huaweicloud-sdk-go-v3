package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginInputVoData struct {

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 租户ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Validation *ExtensionValidation `json:"validation,omitempty"`

	// 样式信息
	LayoutContent *string `json:"layout_content,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginInputVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginInputVoData struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginInputVoData", string(data)}, " ")
}
