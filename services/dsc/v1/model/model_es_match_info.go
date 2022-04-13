package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ES数据项匹配信息
type EsMatchInfo struct {
	// 数据字段名

	FieldName *string `json:"field_name,omitempty"`
	// 规则名

	RuleName *string `json:"rule_name,omitempty"`
	// 规则ID

	RuleId *string `json:"rule_id,omitempty"`
	// 规则风险等级

	RuleRiskLevel *int32 `json:"rule_risk_level,omitempty"`
}

func (o EsMatchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsMatchInfo struct{}"
	}

	return strings.Join([]string{"EsMatchInfo", string(data)}, " ")
}
