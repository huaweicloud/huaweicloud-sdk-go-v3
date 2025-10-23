package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetResourceCopySummaryResponseSummary struct {

	// 统计日期
	StatisticDate *string `json:"statistic_date,omitempty"`

	// 当前统计日期内的备份总数
	TotalCopyCounts *int32 `json:"total_copy_counts,omitempty"`

	// 当前统计日期内的完成备份总数, key -> ResourceCopyStatisticsKeyEnum.COMPLETED.getValue()
	CompletedCopyCounts *int32 `json:"completed_copy_counts,omitempty"`

	// 当前统计日期内的失败备份总数, key -> ResourceCopyStatisticsKeyEnum.FAILED.getValue()
	FailedCopyCounts *int32 `json:"failed_copy_counts,omitempty"`
}

func (o GetResourceCopySummaryResponseSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceCopySummaryResponseSummary struct{}"
	}

	return strings.Join([]string{"GetResourceCopySummaryResponseSummary", string(data)}, " ")
}
