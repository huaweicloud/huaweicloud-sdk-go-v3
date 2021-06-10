package model

import (
	"encoding/json"

	"strings"
)

//
type ContentCompareResult struct {
	// 内容对比的任务id。

	CompareTaskId string `json:"compare_task_id"`
	// 内容对比结果概览。

	ContentCompareOverview *[]ContentCompareResultOverview `json:"content_compare_overview,omitempty"`
	// 内容对比结果概览总数。

	ContentCompareOverviewCount *int32 `json:"content_compare_overview_count,omitempty"`
	// 内容对比结果详情。

	ContentCompareDetails *[]ContentCompareResultDetails `json:"content_compare_details,omitempty"`
	// 内容对比结果差异。

	ContentCompareDiffs *[]ContentCompareResultDiffs `json:"content_compare_diffs,omitempty"`
	// 错误码。

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息。

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ContentCompareResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ContentCompareResult struct{}"
	}

	return strings.Join([]string{"ContentCompareResult", string(data)}, " ")
}
