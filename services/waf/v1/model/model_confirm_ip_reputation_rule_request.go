package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmIpReputationRuleRequest Request Object
type ConfirmIpReputationRuleRequest struct {

	// **参数解释：** policyid **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** ''
	PolicyId string `json:"policy_id"`

	// **参数解释：** geoipRuleId **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** ''
	RuleId string `json:"rule_id"`
}

func (o ConfirmIpReputationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmIpReputationRuleRequest struct{}"
	}

	return strings.Join([]string{"ConfirmIpReputationRuleRequest", string(data)}, " ")
}
