package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDtoInputInfo struct {

	// **参数解释**： 插件输入配置的唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件输入配置的默认值，未填写时默认显示。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： 插件输入配置的输入类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 插件输入配置的输入内容的正则校验规则。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Validation *string `json:"validation,omitempty"`

	// **参数解释**： 样式信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LayoutContent *string `json:"layout_content,omitempty"`
}

func (o PluginDtoInputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoInputInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoInputInfo", string(data)}, " ")
}
