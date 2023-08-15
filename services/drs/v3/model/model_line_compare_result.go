package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LineCompareResult
type LineCompareResult struct {

	// 行对比任务的id。
	CompareTaskId *string `json:"compare_task_id,omitempty"`

	// 行对比结果概览。
	LineCompareOverview *[]LineCompareResultOverview `json:"line_compare_overview,omitempty"`

	// 行对比结果概览总数。
	LineCompareOverviewCount *int32 `json:"line_compare_overview_count,omitempty"`

	// 行对比结果详情。
	LineCompareDetails *[]LineCompareResultDetails `json:"line_compare_details,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o LineCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResult struct{}"
	}

	return strings.Join([]string{"LineCompareResult", string(data)}, " ")
}
