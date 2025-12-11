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

	// 风险规则类型 - LOGIN：登录  - OPERATE：操作
	Type string `json:"type"`

	// 风险规则特征
	Feature *string `json:"feature,omitempty"`

	// 风险规则状态。 - ON: 开启 - OFF: 关闭
	Status string `json:"status"`

	// 风险规则优先级。数字越小优先级越高。
	Rank *int32 `json:"rank,omitempty"`

	// 风险级别 - LOW：低 - MEDIUM：中 - HIGH：高 - NO_RISK：无风险
	RiskLevel *string `json:"risk_level,omitempty"`

	// 规则类型 - SYSTEM: 系统  - CUSTOMIZE: 自定义
	RuleType *string `json:"rule_type,omitempty"`
}

func (o RuleRiskResponseRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRiskResponseRules struct{}"
	}

	return strings.Join([]string{"RuleRiskResponseRules", string(data)}, " ")
}
