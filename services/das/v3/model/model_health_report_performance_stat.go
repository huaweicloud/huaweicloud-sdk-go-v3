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
}

func (o HealthReportPerformanceStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportPerformanceStat struct{}"
	}

	return strings.Join([]string{"HealthReportPerformanceStat", string(data)}, " ")
}
