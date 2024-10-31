package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportRatioStat struct {

	// 指标名。
	Metric string `json:"metric"`

	// 最大值。
	MaxValue float64 `json:"max_value"`

	// 高水位占比。
	CriticalRatio float64 `json:"critical_ratio"`

	// 中水位占比。
	MediumRatio float64 `json:"medium_ratio"`

	// 低水位占比。
	LightRatio float64 `json:"light_ratio"`
}

func (o HealthReportRatioStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportRatioStat struct{}"
	}

	return strings.Join([]string{"HealthReportRatioStat", string(data)}, " ")
}
