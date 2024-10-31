package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceStat struct {

	// 表大小Top列表。
	SizeTop []HealthReportTableSpaceInfo `json:"size_top"`

	// 表行数量Top列表。
	RowsTop []HealthReportTableSpaceInfo `json:"rows_top"`

	// 表大小增长Top列表。
	SizeIncrTop []HealthReportTableSpaceIncrInfo `json:"size_incr_top"`

	// 表行数量增长Top列表。
	RowsIncrTop []HealthReportTableSpaceIncrInfo `json:"rows_incr_top"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportTableSpaceStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceStat struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceStat", string(data)}, " ")
}
