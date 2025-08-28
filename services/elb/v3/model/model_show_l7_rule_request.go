package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowL7RuleRequest Request Object
type ShowL7RuleRequest struct {

	// **参数解释**：7层转发策略。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7policyId string `json:"l7policy_id"`

	// **参数解释**：7层转发规则。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7ruleId string `json:"l7rule_id"`
}

func (o ShowL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7RuleRequest struct{}"
	}

	return strings.Join([]string{"ShowL7RuleRequest", string(data)}, " ")
}
