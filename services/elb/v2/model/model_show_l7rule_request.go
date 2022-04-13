package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowL7ruleRequest struct {
	// 转发策略id

	L7policyId string `json:"l7policy_id"`
	// 转发规则id

	L7ruleId string `json:"l7rule_id"`
}

func (o ShowL7ruleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7ruleRequest struct{}"
	}

	return strings.Join([]string{"ShowL7ruleRequest", string(data)}, " ")
}
