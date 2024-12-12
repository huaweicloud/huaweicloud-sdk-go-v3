package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportPerformanceStat struct {

	// 峰值统计信息列表。
	PeakStats []HealthReportSingleValueStat `json:"peak_stats"`

	// 比率值数据列表。
	RatioStats []HealthReportRatioStat `json:"ratio_stats"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportPerformanceStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportPerformanceStat struct{}"
	}

	return strings.Join([]string{"HealthReportPerformanceStat", string(data)}, " ")
}
