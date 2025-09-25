package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyRuleIdRequestBody struct {

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds *[]PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids,omitempty"`
}

func (o PolicyRuleIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyRuleIdRequestBody struct{}"
	}

	return strings.Join([]string{"PolicyRuleIdRequestBody", string(data)}, " ")
}
