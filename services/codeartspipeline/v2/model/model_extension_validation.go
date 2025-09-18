package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionValidation struct {

	// **参数解释**： 若插件输入项为必填而实际未填写时，提示的信息。 **取值范围**： 不涉及。
	RequiredMessage *string `json:"required_message,omitempty"`

	// **参数解释**： 插件输入项值的校验正则表达式。 **取值范围**： 不涉及。
	Regex *string `json:"regex,omitempty"`

	// **参数解释**： 若插件输入项的值不满足regex中的正则表达式时，提示的信息。 **取值范围**： 不涉及。
	RegexMessage *string `json:"regex_message,omitempty"`

	// **参数解释**： 插件输入项值的最大长度。 **取值范围**： 不涉及。
	MaxLength *int32 `json:"max_length,omitempty"`

	// **参数解释**： 插件输入项值的最小长度。 **取值范围**： 不涉及。
	MinLength *int32 `json:"min_length,omitempty"`
}

func (o ExtensionValidation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionValidation struct{}"
	}

	return strings.Join([]string{"ExtensionValidation", string(data)}, " ")
}
