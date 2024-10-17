package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlRuleResponseRules struct {

	// SQL规则ID
	Id *string `json:"id,omitempty"`

	// SQL规则名称
	Name *string `json:"name,omitempty"`

	// 规则的状态： - ON - OFF
	Status *string `json:"status,omitempty"`

	// 风险级别 - HIGH - MEDIUM - LOW
	RiskLevel *string `json:"risk_level,omitempty"`

	// 风险类型
	Type *string `json:"type,omitempty"`

	// 优先级。数字越小优先级越高。
	Rank *int32 `json:"rank,omitempty"`

	// SQL命令特征
	Feature *string `json:"feature,omitempty"`

	// 正则表达式
	Regex *string `json:"regex,omitempty"`
}

func (o SqlRuleResponseRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlRuleResponseRules struct{}"
	}

	return strings.Join([]string{"SqlRuleResponseRules", string(data)}, " ")
}
