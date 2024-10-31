package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportSummaryInfo struct {

	// 分析结果列表
	AnalysisResults []HealthReportAnalysisResult `json:"analysis_results"`
}

func (o HealthReportSummaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportSummaryInfo struct{}"
	}

	return strings.Join([]string{"HealthReportSummaryInfo", string(data)}, " ")
}
