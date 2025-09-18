package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleProperty struct {

	// **参数解释**： 属性键。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 展示名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 比较运算符。 **约束限制**： 不涉及。 **取值范围**： - ‘=’：等于。 - ‘>’：大于。 - ‘<’：小于。 - ‘>=’：大于等于。 - ‘<=’：小于等于。 - ‘!=’：不等于。 - ‘in’：包含其中。 - ‘not in’：不被包含。 **默认取值**： 不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**： 属性值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 数据类型。 **约束限制**： 不涉及。 **取值范围**： - float：数值型。 - string：字符型。 **默认取值**： 不涉及。
	ValueType string `json:"value_type"`

	// **参数解释**： 属性是否启用。 **约束限制**： 不涉及。 **取值范围**： - true：启用。 - false：未启用。 **默认取值**： false。
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o RuleProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleProperty struct{}"
	}

	return strings.Join([]string{"RuleProperty", string(data)}, " ")
}
