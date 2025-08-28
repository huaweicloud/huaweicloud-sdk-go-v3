package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateL7RuleRequest Request Object
type UpdateL7RuleRequest struct {

	// **参数解释**：策略ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7policyId string `json:"l7policy_id"`

	// **参数解释**：规则ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7ruleId string `json:"l7rule_id"`

	Body *UpdateL7RuleRequestBody `json:"body,omitempty"`
}

func (o UpdateL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequest", string(data)}, " ")
}
