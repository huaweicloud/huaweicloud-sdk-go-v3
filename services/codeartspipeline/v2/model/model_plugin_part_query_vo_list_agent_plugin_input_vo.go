package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginInputVo struct {

	// **参数解释**： 扩展插件名称。 **取值范围**： 1到50位字符。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件展示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 插件输入项详细信息。 **取值范围**： 不涉及。
	Data *[]PluginPartQueryVoListAgentPluginInputVoData `json:"data,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginInputVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginInputVo struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginInputVo", string(data)}, " ")
}
