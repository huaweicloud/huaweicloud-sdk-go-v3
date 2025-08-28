package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteL7RuleRequest Request Object
type DeleteL7RuleRequest struct {

	// **参数解释**：策略ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7policyId string `json:"l7policy_id"`

	// **参数解释**：规则ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7ruleId string `json:"l7rule_id"`
}

func (o DeleteL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7RuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleRequest", string(data)}, " ")
}
