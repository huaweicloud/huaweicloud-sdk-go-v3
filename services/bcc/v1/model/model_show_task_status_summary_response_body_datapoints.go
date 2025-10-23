package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTaskStatusSummaryResponseBodyDatapoints struct {

	// 统计日期
	StatisticsDate *string `json:"statistics_date,omitempty"`

	// 当前统计周期内的总任务数
	TotalTaskCount *int32 `json:"total_task_count,omitempty"`

	// 当前统计周期内的已完成的任务数
	CompletedTaskCount *int32 `json:"completed_task_count,omitempty"`

	// 当前统计周期内的失败任务数
	FailedTaskCount *int32 `json:"failed_task_count,omitempty"`

	// 当前统计周期内的运行中任务数
	RunningTaskCount *int32 `json:"running_task_count,omitempty"`

	// 当前统计周期内的跳过任务数
	SkippedTaskCount *int32 `json:"skipped_task_count,omitempty"`

	// 当前统计周期内的超时任务数
	TimeoutTaskCount *int32 `json:"timeout_task_count,omitempty"`
}

func (o ShowTaskStatusSummaryResponseBodyDatapoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatusSummaryResponseBodyDatapoints struct{}"
	}

	return strings.Join([]string{"ShowTaskStatusSummaryResponseBodyDatapoints", string(data)}, " ")
}
