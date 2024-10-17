package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlRuleRequest struct {

	// 风险级别 - HIGH - MEDIUM - LOW - NO_RISK
	RiskLevels *string `json:"risk_levels,omitempty"`
}

func (o SqlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlRuleRequest struct{}"
	}

	return strings.Join([]string{"SqlRuleRequest", string(data)}, " ")
}
