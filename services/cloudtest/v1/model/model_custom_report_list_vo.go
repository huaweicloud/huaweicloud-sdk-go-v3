package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实际的数据类型：单个对象，集合或NULL
type CustomReportListVo struct {

	// 报表id
	Id *string `json:"id,omitempty"`

	// 报表名称
	Name *string `json:"name,omitempty"`

	Filter *ReportFilter `json:"filter,omitempty"`

	// 工件类型(用例：case,测试套：task)
	WorkpieceType *string `json:"workpiece_type,omitempty"`

	// 分析维度
	AnalysisDimension *string `json:"analysis_dimension,omitempty"`

	// 对比维度
	CompareDimension *string `json:"compare_dimension,omitempty"`

	// 报表数据
	ChartData *[]ReportChartDataVo `json:"chart_data,omitempty"`
}

func (o CustomReportListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomReportListVo struct{}"
	}

	return strings.Join([]string{"CustomReportListVo", string(data)}, " ")
}
