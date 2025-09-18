package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionExtendProp struct {

	// **参数解释**： API选项。 **取值范围**： 不涉及。
	ApiOptions *string `json:"api_options,omitempty"`

	// **参数解释**： API类型。 **取值范围**： 不涉及。
	ApiType *string `json:"api_type,omitempty"`

	// **参数解释**： 显示占位符。 **取值范围**： 不涉及。
	ShowPlaceholder *string `json:"show_placeholder,omitempty"`

	// **参数解释**： 选项。 **取值范围**： 不涉及。
	Options *string `json:"options,omitempty"`

	// **参数解释**： 禁用条件。 **取值范围**： 不涉及。
	DisabledConditions *string `json:"disabled_conditions,omitempty"`

	// **参数解释**： 可见条件。 **取值范围**： 不涉及。
	VisibleConditions *string `json:"visible_conditions,omitempty"`
}

func (o ExtensionExtendProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionExtendProp struct{}"
	}

	return strings.Join([]string{"ExtensionExtendProp", string(data)}, " ")
}
