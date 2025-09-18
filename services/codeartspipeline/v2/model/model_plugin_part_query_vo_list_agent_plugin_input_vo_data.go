package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryVoListAgentPluginInputVoData struct {

	// **参数解释**： 插件输入项唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 插件输入项名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件输入项默认值。 **取值范围**： 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： 插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 插件版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件输入项类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 租户id。 **取值范围**： 32位字符，由数字和字母组成。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Validation *ExtensionValidation `json:"validation,omitempty"`

	// **参数解释**： 插件输入项样式信息。 **取值范围**： 不涉及。
	LayoutContent *string `json:"layout_content,omitempty"`

	ExtendProp *ExtensionExtendProp `json:"extend_prop,omitempty"`

	// **参数解释**： 前端渲染使用的数据信息。 **取值范围**： 不涉及。
	FrontDataProp *string `json:"front_data_prop,omitempty"`

	// **参数解释**： 标签。 **取值范围**： 不涉及。
	Label *string `json:"label,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否必须。 **取值范围**： 不涉及。
	Required *string `json:"required,omitempty"`
}

func (o PluginPartQueryVoListAgentPluginInputVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryVoListAgentPluginInputVoData struct{}"
	}

	return strings.Join([]string{"PluginPartQueryVoListAgentPluginInputVoData", string(data)}, " ")
}
