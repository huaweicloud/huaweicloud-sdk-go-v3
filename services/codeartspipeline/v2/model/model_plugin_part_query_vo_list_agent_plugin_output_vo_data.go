package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginOutputVoData struct {

	// **参数解释**： 唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 扩展插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**： 定义的插件输出名称。 **取值范围**： 不涉及。
	OutputKey *string `json:"output_key,omitempty"`

	// **参数解释**： 定义的插件输出内容。 **取值范围**： 不涉及。
	OutputValue *string `json:"output_value,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginOutputVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginOutputVoData struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginOutputVoData", string(data)}, " ")
}
