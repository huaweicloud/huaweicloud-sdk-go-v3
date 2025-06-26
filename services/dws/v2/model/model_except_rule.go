package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptRule **参数解释**： 异常规则信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ExceptRule struct {

	// **参数解释**： 规则项名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RuleKey *string `json:"rule_key,omitempty"`

	// **参数解释**： 规则项数值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RuleValue *string `json:"rule_value,omitempty"`
}

func (o ExceptRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptRule struct{}"
	}

	return strings.Join([]string{"ExceptRule", string(data)}, " ")
}
