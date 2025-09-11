package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleRiskStatisticsDto struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 实例数据库风险汇总
	InstanceRiskCount *[]AuditInsanceRiskCount `json:"instance_risk_count,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`
}

func (o RuleRiskStatisticsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRiskStatisticsDto struct{}"
	}

	return strings.Join([]string{"RuleRiskStatisticsDto", string(data)}, " ")
}
