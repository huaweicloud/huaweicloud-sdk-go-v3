package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleContent struct {

	// **参数解释**： 分组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName string `json:"group_name"`

	// **参数解释**： 继承后的子策略是否可以修改阈值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CanModifyWhenInherit *bool `json:"can_modify_when_inherit,omitempty"`

	// **参数解释**： 规则属性列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Properties []RuleProperty `json:"properties"`
}

func (o RuleContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleContent struct{}"
	}

	return strings.Join([]string{"RuleContent", string(data)}, " ")
}
