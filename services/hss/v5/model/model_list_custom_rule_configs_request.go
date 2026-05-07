package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRuleConfigsRequest Request Object
type ListCustomRuleConfigsRequest struct {

	// **参数解释**： 规则ID **约束限制**： 不涉及 **取值范围**： 字符长度1-36位 **默认取值**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 规则名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomRuleConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRuleConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomRuleConfigsRequest", string(data)}, " ")
}
