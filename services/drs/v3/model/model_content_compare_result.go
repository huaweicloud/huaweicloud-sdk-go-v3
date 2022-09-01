package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ContentCompareResult struct {

	// 内容对比的任务id。
	CompareTaskId string `json:"compare_task_id" xml:"compare_task_id"`

	// 内容对比结果概览。
	ContentCompareOverview *[]ContentCompareResultOverview `json:"content_compare_overview,omitempty" xml:"content_compare_overview"`

	// 内容对比结果概览总数。
	ContentCompareOverviewCount *int32 `json:"content_compare_overview_count,omitempty" xml:"content_compare_overview_count"`

	// 内容对比结果详情。
	ContentCompareDetails *[]ContentCompareResultDetails `json:"content_compare_details,omitempty" xml:"content_compare_details"`

	// 内容对比结果差异。
	ContentCompareDiffs *[]ContentCompareResultDiffs `json:"content_compare_diffs,omitempty" xml:"content_compare_diffs"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o ContentCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResult struct{}"
	}

	return strings.Join([]string{"ContentCompareResult", string(data)}, " ")
}
