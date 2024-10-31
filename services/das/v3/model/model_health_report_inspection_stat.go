package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportInspectionStat struct {

	// 巡检评分。
	InspectionScore []HealthReportInspectionScore `json:"inspection_score"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportInspectionStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportInspectionStat struct{}"
	}

	return strings.Join([]string{"HealthReportInspectionStat", string(data)}, " ")
}
