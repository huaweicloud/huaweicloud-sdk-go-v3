package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginInstanceVoInputInfo struct {

	// **参数解释**： 插件输入项唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 插件输入项名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件输入项默认值。 **取值范围**： 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： 扩展插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件输入项类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 租户id。 **取值范围**： 32位字符，由数字和字母组成。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Validation *ExtensionValidation `json:"validation,omitempty"`

	// **参数解释**： 插件输入项样式信息。 **取值范围**： 不涉及。
	LayoutContent *string `json:"layout_content,omitempty"`
}

func (o PluginInstanceVoInputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginInstanceVoInputInfo struct{}"
	}

	return strings.Join([]string{"PluginInstanceVoInputInfo", string(data)}, " ")
}
