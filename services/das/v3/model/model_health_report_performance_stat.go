package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportPerformanceStat struct {

	// 峰值统计信息列表。
	PeakStats *[]HealthReportSingleValueStat `json:"peak_stats,omitempty"`

	// 最新统计信息列表。
	LastValues *[]HealthReportSingleValueStat `json:"last_values,omitempty"`

	// 比率值数据列表。
	RatioStats *[]HealthReportRatioStat `json:"ratio_stats,omitempty"`

	// 统计分析是否成功。
	AnalyzeSuccess *bool `json:"analyze_success,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o HealthReportPerformanceStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportPerformanceStat struct{}"
	}

	return strings.Join([]string{"HealthReportPerformanceStat", string(data)}, " ")
}
