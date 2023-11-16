package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginInputVo struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 结果集
	Data *[]PluginPartQueryVoListAgentPluginInputVoData `json:"data,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginInputVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginInputVo struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginInputVo", string(data)}, " ")
}
