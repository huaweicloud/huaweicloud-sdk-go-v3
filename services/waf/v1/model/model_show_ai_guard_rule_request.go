package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiGuardRuleRequest Request Object
type ShowAiGuardRuleRequest struct {

	// **参数解释：** 策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 规则ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o ShowAiGuardRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAiGuardRuleRequest", string(data)}, " ")
}
