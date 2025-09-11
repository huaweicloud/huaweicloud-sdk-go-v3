package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskRuleStatistic struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则英文名
	RuleNameEnglish *string `json:"rule_name_english,omitempty"`
}

func (o RiskRuleStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskRuleStatistic struct{}"
	}

	return strings.Join([]string{"RiskRuleStatistic", string(data)}, " ")
}
