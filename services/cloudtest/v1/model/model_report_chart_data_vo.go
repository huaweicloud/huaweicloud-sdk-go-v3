package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportChartDataVo 报表数据
type ReportChartDataVo struct {

	// 报表id
	Id *string `json:"id,omitempty"`

	// 报表名称
	Name *string `json:"name,omitempty"`

	AnalyzeDim *ReportDimVo `json:"analyze_dim,omitempty"`

	// 对比维度数据
	CompareDim *[]ReportDimVo `json:"compare_dim,omitempty"`
}

func (o ReportChartDataVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportChartDataVo struct{}"
	}

	return strings.Join([]string{"ReportChartDataVo", string(data)}, " ")
}
