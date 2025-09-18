package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginBasicDto struct {

	// **参数解释**： 扩展插件UUID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 扩展插件图标URL。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 运行属性。 **约束限制**： 不涉及。 **取值范围**： - agent：基于agent运行。 - agentless：无需agent运行。 **默认取值**： 不涉及。
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// **参数解释**： 扩展插件唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginName string `json:"plugin_name"`

	// **参数解释**： 扩展插件展示名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DisplayName string `json:"display_name"`

	// **参数解释**： 扩展插件业务类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BusinessType string `json:"business_type"`

	// **参数解释**： 扩展插件业务类型展示名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BusinessTypeDisplayName string `json:"business_type_display_name"`

	// **参数解释**： 扩展插件描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否私有。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IsPrivate *int32 `json:"is_private,omitempty"`

	// **参数解释**： 局点。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 维护者。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Maintainers *string `json:"maintainers,omitempty"`

	// **参数解释**： 标识是否为多个step组成的组。 **约束限制**： 不涉及。 **取值范围**： - single：单step插件。 - multi：组合插件。 **默认取值**： 不涉及。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`
}

func (o PluginBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginBasicDto struct{}"
	}

	return strings.Join([]string{"PluginBasicDto", string(data)}, " ")
}
