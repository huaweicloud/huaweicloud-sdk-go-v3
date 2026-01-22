package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderRuleAclDto **参数解释**： 修改规则顺序请求体 **约束限制**： 不涉及
type OrderRuleAclDto struct {

	// **参数解释**： 目标规则ID，填写后添加的新规则置位于此规则之后，非置顶时不能为空，置顶时为空，目标规则ID可以通过[查询防护规则接口](ListAclRules.xml)获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。 **约束限制**： 非置顶时目标规则需要存在 **取值范围**： 不涉及 **默认取值**： 不涉及
	DestRuleId *string `json:"dest_rule_id,omitempty"`

	// **参数解释**： 规则是否置顶，用于确认规则是否置顶 **约束限制**： 不涉及 **取值范围**： 0代表非置顶，1代表置顶 **默认取值**： 不涉及
	Top *int32 `json:"top,omitempty"`

	// **参数解释**： 规则是否置底，用于确认规则是否置底 **约束限制**： 不涉及 **取值范围**： 0代表非置底，1代表置底 **默认取值**： 不涉及
	Bottom *int32 `json:"bottom,omitempty"`
}

func (o OrderRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderRuleAclDto struct{}"
	}

	return strings.Join([]string{"OrderRuleAclDto", string(data)}, " ")
}
