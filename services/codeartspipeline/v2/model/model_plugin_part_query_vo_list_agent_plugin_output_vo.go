package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginOutputVo struct {

	// **参数解释**： 扩展插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件展示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 扩展插件输出详细信息。 **取值范围**： 不涉及。
	Data *[]PluginPartQueryVoListAgentPluginOutputVoData `json:"data,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginOutputVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginOutputVo struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginOutputVo", string(data)}, " ")
}
