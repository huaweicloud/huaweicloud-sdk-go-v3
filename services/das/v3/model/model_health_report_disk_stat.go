package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportDiskStat struct {

	// 峰值统计信息列表。
	PeakStats *[]HealthReportSingleValueStat `json:"peak_stats,omitempty"`

	// 比率值数据列表。
	RatioStats *[]HealthReportRatioStat `json:"ratio_stats,omitempty"`

	// 最新统计信息列表。
	LastStat *[]HealthReportSingleValueStat `json:"last_stat,omitempty"`

	// 统计分析是否成功。
	AnalyzeSuccess *bool `json:"analyze_success,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o HealthReportDiskStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportDiskStat struct{}"
	}

	return strings.Join([]string{"HealthReportDiskStat", string(data)}, " ")
}
