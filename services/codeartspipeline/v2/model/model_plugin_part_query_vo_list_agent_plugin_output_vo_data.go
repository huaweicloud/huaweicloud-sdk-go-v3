package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginOutputVoData struct {

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 租户ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 名称
	OutputKey *string `json:"output_key,omitempty"`

	// 值
	OutputValue *string `json:"output_value,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginOutputVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginOutputVoData struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginOutputVoData", string(data)}, " ")
}
