package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportDiskStat struct {

	// 峰值统计信息列表。
	PeakStats []HealthReportSingleValueStat `json:"peak_stats"`

	// 比率值数据列表。
	RatioStats []HealthReportRatioStat `json:"ratio_stats"`
}

func (o HealthReportDiskStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportDiskStat struct{}"
	}

	return strings.Join([]string{"HealthReportDiskStat", string(data)}, " ")
}
