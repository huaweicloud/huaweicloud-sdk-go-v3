package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleRiskResponseRules struct {

	// 风险规则ID
	Id string `json:"id"`

	// 风险规则名称
	Name string `json:"name"`

	// 风险规则类型
	Type string `json:"type"`

	// 风险规则特征
	Feature *string `json:"feature,omitempty"`

	// 风险规则状态。 - ON: 开启 - OFF: 关闭
	Status string `json:"status"`

	// 风险规则优先级。数字越小优先级越高。
	Rank *int32 `json:"rank,omitempty"`

	// 风险级别 - LOW - MEDIUM - HIGH - NO_RISK]
	RiskLevel *string `json:"risk_level,omitempty"`

	// 规则类型
	RuleType *string `json:"rule_type,omitempty"`
}

func (o RuleRiskResponseRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRiskResponseRules struct{}"
	}

	return strings.Join([]string{"RuleRiskResponseRules", string(data)}, " ")
}
