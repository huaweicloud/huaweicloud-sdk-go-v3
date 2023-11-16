package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginOutputVo struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 结果集
	Data *[]PluginPartQueryVoListAgentPluginOutputVoData `json:"data,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginOutputVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginOutputVo struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginOutputVo", string(data)}, " ")
}
