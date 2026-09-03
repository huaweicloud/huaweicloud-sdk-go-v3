package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableSpaceStat 表空间统计分析
type TableSpaceStat struct {

	// 表大小Top列表
	SizeTop *[]HealthReportTableSpaceInfo `json:"size_top,omitempty"`

	// 表行数量Top列表
	RowsTop *[]HealthReportTableSpaceInfo `json:"rows_top,omitempty"`

	// 表大小增长Top列表
	SizeIncrTop *[]HealthReportTableSpaceIncrInfo `json:"size_incr_top,omitempty"`

	// 表行数量增长Top列表
	RowsIncrTop *[]HealthReportTableSpaceIncrInfo `json:"rows_incr_top,omitempty"`

	// 统计分析是否成功
	AnalyzeSuccess *bool `json:"analyze_success,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o TableSpaceStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableSpaceStat struct{}"
	}

	return strings.Join([]string{"TableSpaceStat", string(data)}, " ")
}
