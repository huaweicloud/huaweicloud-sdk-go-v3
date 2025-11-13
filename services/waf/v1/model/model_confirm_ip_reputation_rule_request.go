package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmIpReputationRuleRequest Request Object
type ConfirmIpReputationRuleRequest struct {

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id，响应体的id字段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 威胁情报规则ip，从接口 “查询威胁情报规则列表”（ListIpReputationRules）获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o ConfirmIpReputationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmIpReputationRuleRequest struct{}"
	}

	return strings.Join([]string{"ConfirmIpReputationRuleRequest", string(data)}, " ")
}
