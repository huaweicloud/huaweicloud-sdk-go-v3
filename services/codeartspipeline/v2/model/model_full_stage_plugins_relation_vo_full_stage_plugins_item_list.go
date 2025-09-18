package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoFullStagePluginsItemList struct {

	// **参数解释**： 插件列表。 **取值范围**： 不涉及。
	PluginsList *[]FullStagePluginsRelationVoPluginsList `json:"plugins_list,omitempty"`

	// **参数解释**： 扩展插件展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 扩展插件业务类型。 **取值范围**： 不涉及。
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**： 扩展插件唯一ID。 **默认取值**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 运行条件。 **取值范围**： 不涉及。
	Conditions *[]string `json:"conditions,omitempty"`

	// **参数解释**： 额外属性。 **取值范围**： 不涉及。
	Addables *[]FullStagePluginsRelationVoAddables `json:"addables,omitempty"`

	// **参数解释**： 是否可编辑。 **取值范围**： - true：可编辑。 - false：不可编辑。
	Editable *bool `json:"editable,omitempty"`

	// **参数解释**： 是否可移除。 **取值范围**： - true：可移除。 - false：不可移除。
	Removable *bool `json:"removable,omitempty"`

	// **参数解释**： 是否可复制。 **取值范围**： - true：可复制。 - false：不可复制。
	Cloneable *bool `json:"cloneable,omitempty"`

	// **参数解释**： 是否禁用。 **取值范围**： - true：禁用。 - false：未禁用。
	Disabled *bool `json:"disabled,omitempty"`
}

func (o FullStagePluginsRelationVoFullStagePluginsItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoFullStagePluginsItemList struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoFullStagePluginsItemList", string(data)}, " ")
}
