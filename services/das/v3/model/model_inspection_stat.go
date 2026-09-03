package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InspectionStat 巡检评分统计分析
type InspectionStat struct {

	// 巡检评分
	InspectionScore *[]HealthReportInspectionScore `json:"inspection_score,omitempty"`

	// 统计分析是否成功
	AnalyzeSuccess *bool `json:"analyze_success,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o InspectionStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InspectionStat struct{}"
	}

	return strings.Join([]string{"InspectionStat", string(data)}, " ")
}
