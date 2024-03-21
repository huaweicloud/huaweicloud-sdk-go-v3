package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenerateReportInfo struct {

	// 报表名称
	Name *string `json:"name,omitempty"`

	// 工件类型(用例：case,测试套：suite)
	WorkpieceType *string `json:"workpiece_type,omitempty"`

	// 分析维度。该参数选择横坐标（X轴）维度，不传会返回空列表。
	AnalysisDimRow *string `json:"analysis_dim_row,omitempty"`

	// 对比维度
	CompareDimColumn *string `json:"compare_dim_column,omitempty"`

	Filter *ReportFilter `json:"filter,omitempty"`
}

func (o GenerateReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateReportInfo struct{}"
	}

	return strings.Join([]string{"GenerateReportInfo", string(data)}, " ")
}
