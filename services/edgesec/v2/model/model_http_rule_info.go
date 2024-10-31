package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpRuleInfo Http防护策略和规则信息
type HttpRuleInfo struct {

	// 策略下的规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 策略下的规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o HttpRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpRuleInfo struct{}"
	}

	return strings.Join([]string{"HttpRuleInfo", string(data)}, " ")
}
