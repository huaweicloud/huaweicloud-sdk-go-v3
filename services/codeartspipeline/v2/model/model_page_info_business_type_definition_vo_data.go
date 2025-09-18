package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoBusinessTypeDefinitionVoData struct {

	// **参数解释**： 业务类型。 **取值范围**： 不涉及。
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**： 展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 是否可编辑。 **取值范围**： - true：可编辑。 - false：不可编辑。
	Editable *bool `json:"editable,omitempty"`

	// **参数解释**： 是否可移除。 **取值范围**： - true：可移除。 - false：不可移除。
	Removable *bool `json:"removable,omitempty"`

	// **参数解释**： 是否可复制。 **取值范围**： - true：可复制。 - false：不可复制。
	Cloneable *bool `json:"cloneable,omitempty"`

	// **参数解释**： 是否禁用。 **取值范围**： - true：禁用。 - false：未禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// **参数解释**： 是否可添加。 **取值范围**： 不涉及。
	Addables *[]map[string]bool `json:"addables,omitempty"`

	// **参数解释**： 条件。 **取值范围**： 不涉及。
	Conditions *[]string `json:"conditions,omitempty"`

	// **参数解释**： 插件列表。 **取值范围**： 不涉及。
	PluginsList *[]PageInfoBusinessTypeDefinitionVoPluginsList `json:"plugins_list,omitempty"`
}

func (o PageInfoBusinessTypeDefinitionVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoBusinessTypeDefinitionVoData struct{}"
	}

	return strings.Join([]string{"PageInfoBusinessTypeDefinitionVoData", string(data)}, " ")
}
