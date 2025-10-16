package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCustomRulesResponse Response Object
type BatchUpdateCustomRulesResponse struct {

	// **参数解释：** 策略和规则id数组，返回防护策略与对应规则的ID关联关系 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds  *[]PolicyRuleIdResponseBodyPolicyRuleIds `json:"policy_rule_ids,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o BatchUpdateCustomRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCustomRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateCustomRulesResponse", string(data)}, " ")
}
