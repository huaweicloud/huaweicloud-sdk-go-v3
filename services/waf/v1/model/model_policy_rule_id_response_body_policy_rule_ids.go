package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyRuleIdResponseBodyPolicyRuleIds struct {

	// **参数解释：** 策略id，唯一标识一条防护策略 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释：** 规则id数组，返回当前防护策略下的多条规则ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleIds *[]string `json:"rule_ids,omitempty"`
}

func (o PolicyRuleIdResponseBodyPolicyRuleIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyRuleIdResponseBodyPolicyRuleIds struct{}"
	}

	return strings.Join([]string{"PolicyRuleIdResponseBodyPolicyRuleIds", string(data)}, " ")
}
