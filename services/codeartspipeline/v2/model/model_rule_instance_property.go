package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstanceProperty struct {

	// **参数解释**： 规则属性键。 **取值范围**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 规则类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 展示名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 比较运算符。 **取值范围**： 不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**： 属性值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 数据类型。 **取值范围**： 不涉及。
	ValueType string `json:"value_type"`
}

func (o RuleInstanceProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstanceProperty struct{}"
	}

	return strings.Join([]string{"RuleInstanceProperty", string(data)}, " ")
}
