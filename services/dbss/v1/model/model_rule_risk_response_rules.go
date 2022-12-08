package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleRiskResponseRules struct {

	// 风险规则ID
	Id *string `json:"id,omitempty"`

	// 风险规则名称
	Name *string `json:"name,omitempty"`

	// 风险类型
	Type *string `json:"type,omitempty"`

	// 风险特征
	Feature *string `json:"feature,omitempty"`

	// 风险规则状态
	Status *string `json:"status,omitempty"`

	// 风险规则优先级
	Rank *int32 `json:"rank,omitempty"`

	// 风险级别
	RiskLevel *string `json:"risk_level,omitempty"`
}

func (o RuleRiskResponseRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRiskResponseRules struct{}"
	}

	return strings.Join([]string{"RuleRiskResponseRules", string(data)}, " ")
}
