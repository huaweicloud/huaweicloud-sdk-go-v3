package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SummaryInfo 日报内容摘要信息
type SummaryInfo struct {

	// 分析结果列表
	AnalysisResults *[]HealthReportAnalysisResult `json:"analysis_results,omitempty"`

	// 健康等级
	HealthRank *string `json:"health_rank,omitempty"`
}

func (o SummaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummaryInfo struct{}"
	}

	return strings.Join([]string{"SummaryInfo", string(data)}, " ")
}
