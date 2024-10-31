package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportAnalysisResult struct {

	// 风险点编码。
	RiskCode string `json:"risk_code"`

	// 风险点级别。
	RiskLevel string `json:"risk_level"`

	// 风险点内容。
	RiskContent string `json:"risk_content"`

	// 可能原因列表。
	Reasons []HealthReportRiskReason `json:"reasons"`
}

func (o HealthReportAnalysisResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportAnalysisResult struct{}"
	}

	return strings.Join([]string{"HealthReportAnalysisResult", string(data)}, " ")
}
