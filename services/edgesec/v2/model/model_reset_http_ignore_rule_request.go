package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetHttpIgnoreRuleRequest Request Object
type ResetHttpIgnoreRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ResetHttpIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetHttpIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"ResetHttpIgnoreRuleRequest", string(data)}, " ")
}
