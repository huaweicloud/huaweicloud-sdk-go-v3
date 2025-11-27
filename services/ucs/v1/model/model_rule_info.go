package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInfo struct {

	// 权限策略id
	RuleID *string `json:"ruleID,omitempty"`

	// 权限策略名字
	RuleName *string `json:"ruleName,omitempty"`
}

func (o RuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInfo struct{}"
	}

	return strings.Join([]string{"RuleInfo", string(data)}, " ")
}
