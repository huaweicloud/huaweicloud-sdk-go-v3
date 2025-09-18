package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewExtensionInputs struct {

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 标签。 **取值范围**： 不涉及。
	Label *string `json:"label,omitempty"`

	// **参数解释**： 说明。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 默认值。 **取值范围**： 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： 是否必填。 **取值范围**： - true：必填。 - false：非必填。
	Required *bool `json:"required,omitempty"`

	ExtendProp *ExtensionExtendProp `json:"extend_prop,omitempty"`

	Validation *ExtensionValidation `json:"validation,omitempty"`
}

func (o NewExtensionInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionInputs struct{}"
	}

	return strings.Join([]string{"NewExtensionInputs", string(data)}, " ")
}
