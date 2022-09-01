package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbMatchInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty" xml:"column_name"`

	// 匹配的规则名
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name"`

	// 匹配的规则ID
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 匹配规则风险等级
	RuleRiskLevel *int32 `json:"rule_risk_level,omitempty" xml:"rule_risk_level"`

	// 风险数据行
	ColumnLine *[]int64 `json:"column_line,omitempty" xml:"column_line"`
}

func (o DbMatchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbMatchInfo struct{}"
	}

	return strings.Join([]string{"DbMatchInfo", string(data)}, " ")
}
