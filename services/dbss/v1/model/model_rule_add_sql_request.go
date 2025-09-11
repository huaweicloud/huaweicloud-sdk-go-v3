package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleAddSqlRequest struct {

	// 风险等级  - LOW：低  - MEDIUM：中  - HIGH：高 - NO_RISK：无
	RiskLevel string `json:"risk_level"`

	// 规则名称
	RuleName string `json:"rule_name"`

	// 正则表达式
	SqlRegex string `json:"sql_regex"`

	// 状态  - ON：开启  - OFF：关闭
	Status string `json:"status"`
}

func (o RuleAddSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddSqlRequest struct{}"
	}

	return strings.Join([]string{"RuleAddSqlRequest", string(data)}, " ")
}
