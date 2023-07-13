package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleInfo 使用该Ip地址相关的规则信息
type RuleInfo struct {

	// 规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o RuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInfo struct{}"
	}

	return strings.Join([]string{"RuleInfo", string(data)}, " ")
}
