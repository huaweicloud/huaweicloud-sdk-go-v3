package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstanceContent struct {

	// **参数解释**： 分组名称。 **取值范围**： 不涉及。
	GroupName string `json:"group_name"`

	// **参数解释**： 是否可编辑。 **取值范围**： 不涉及。
	Editable *bool `json:"editable,omitempty"`

	// **参数解释**： 分组类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 继承后的子策略是否可以修改阈值。 **取值范围**： - true：可以修改阈值。 - false：不可以修改阈值。
	CanModifyWhenInherit *bool `json:"can_modify_when_inherit,omitempty"`

	// **参数解释**： 规则属性列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Properties []RuleInstanceProperty `json:"properties"`
}

func (o RuleInstanceContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstanceContent struct{}"
	}

	return strings.Join([]string{"RuleInstanceContent", string(data)}, " ")
}
