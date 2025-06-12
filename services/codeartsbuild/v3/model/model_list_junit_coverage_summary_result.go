package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJunitCoverageSummaryResult 返回结果
type ListJunitCoverageSummaryResult struct {

	// 单元测试覆盖率报告列表
	UnitSummaryList *[]ListJunitCoverageSummaryResultUnitSummaryList `json:"unit_summary_list,omitempty"`
}

func (o ListJunitCoverageSummaryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJunitCoverageSummaryResult struct{}"
	}

	return strings.Join([]string{"ListJunitCoverageSummaryResult", string(data)}, " ")
}
