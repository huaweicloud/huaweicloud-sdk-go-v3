package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OprReportInfo struct {

	// 报表名称
	Name *string `json:"name,omitempty"`

	// 报表类型 1：首页用例库， 2：质量报告
	Type *int32 `json:"type,omitempty"`

	// 工件类型(用例：case,测试套：suite)
	WorkpieceType *string `json:"workpiece_type,omitempty"`

	// 分析维度
	AnalysisDimRow *string `json:"analysis_dim_row,omitempty"`

	// 对比维度
	CompareDimColumn *string `json:"compare_dim_column,omitempty"`

	Filter *ReportFilter `json:"filter,omitempty"`
}

func (o OprReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OprReportInfo struct{}"
	}

	return strings.Join([]string{"OprReportInfo", string(data)}, " ")
}
