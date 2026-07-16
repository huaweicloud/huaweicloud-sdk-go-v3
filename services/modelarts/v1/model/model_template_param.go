package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateParam 模板参数params
type TemplateParam struct {

	// **参数解释**：参数名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：参数描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：参数取值。 **取值范围**：不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**：是否展示在console。 **取值范围**：- true   -false。
	Visible *string `json:"visible,omitempty"`

	// **参数解释**：正则校验。 **取值范围**：不涉及。
	Regex *string `json:"regex,omitempty"`
}

func (o TemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateParam struct{}"
	}

	return strings.Join([]string{"TemplateParam", string(data)}, " ")
}
