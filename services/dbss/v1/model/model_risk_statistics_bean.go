package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskStatisticsBean struct {

	// 高风险数量
	HighRiskNum *int64 `json:"high_risk_num,omitempty"`

	// 低风险数量
	LowRiskNum *int64 `json:"low_risk_num,omitempty"`

	// 中风险数量
	MiddleRiskNum *int64 `json:"middle_risk_num,omitempty"`

	// 无风险数量
	NoRiskNum *int64 `json:"no_risk_num,omitempty"`

	// 周期
	Period *string `json:"period,omitempty"`
}

func (o RiskStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskStatisticsBean struct{}"
	}

	return strings.Join([]string{"RiskStatisticsBean", string(data)}, " ")
}
