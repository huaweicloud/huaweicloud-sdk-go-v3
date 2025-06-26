package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExceptRuleReq **参数解释**： 创建/更新异常规则请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type AddExceptRuleReq struct {

	// **参数解释**： 异常规则名称。 **约束限制**： 名称不能为空。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**： 异常规则配置项。 **约束限制**： 不能为空。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ExceptRules *[]ExceptRule `json:"except_rules,omitempty"`
}

func (o AddExceptRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExceptRuleReq struct{}"
	}

	return strings.Join([]string{"AddExceptRuleReq", string(data)}, " ")
}
