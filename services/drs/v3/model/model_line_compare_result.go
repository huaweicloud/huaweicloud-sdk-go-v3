package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LineCompareResult struct {

	// 行对比任务的id。
	CompareTaskId *string `json:"compare_task_id,omitempty" xml:"compare_task_id"`

	// 行对比结果概览。
	LineCompareOverview *[]LineCompareResultOverview `json:"line_compare_overview,omitempty" xml:"line_compare_overview"`

	// 行对比结果概览总数。
	LineCompareOverviewCount *int32 `json:"line_compare_overview_count,omitempty" xml:"line_compare_overview_count"`

	// 行对比结果详情。
	LineCompareDetails *[]LineCompareResultDetails `json:"line_compare_details,omitempty" xml:"line_compare_details"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o LineCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResult struct{}"
	}

	return strings.Join([]string{"LineCompareResult", string(data)}, " ")
}
