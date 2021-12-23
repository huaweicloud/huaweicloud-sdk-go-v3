package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowL7RuleRequest struct {
	// 7层转发策略。

	L7policyId string `json:"l7policy_id"`
	// 7层转发规则。

	L7ruleId string `json:"l7rule_id"`
}

func (o ShowL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7RuleRequest struct{}"
	}

	return strings.Join([]string{"ShowL7RuleRequest", string(data)}, " ")
}
