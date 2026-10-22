package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardRuleRequest Request Object
type UpdateAiGuardRuleRequest struct {

	// **参数解释：** 策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 规则ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`

	Body *UpdateAiGuardRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateAiGuardRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardRuleRequest", string(data)}, " ")
}
