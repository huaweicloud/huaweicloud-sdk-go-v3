package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskStatusSummaryResponseBodySummary 统计数据
type ShowTaskStatusSummaryResponseBodySummary struct {

	// 总任务数据
	TotalTaskCount int32 `json:"total_task_count"`

	// 已完成的总任务数
	CompletedTaskCount int32 `json:"completed_task_count"`

	// 失败的总任务数
	FailedTaskCount int32 `json:"failed_task_count"`

	// 运行中的总任务数
	RunningTaskCount int32 `json:"running_task_count"`

	// 跳过的总任务数
	SkippedTaskCount int32 `json:"skipped_task_count"`

	// 超时的总任务数
	TimeoutTaskCount int32 `json:"timeout_task_count"`
}

func (o ShowTaskStatusSummaryResponseBodySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatusSummaryResponseBodySummary struct{}"
	}

	return strings.Join([]string{"ShowTaskStatusSummaryResponseBodySummary", string(data)}, " ")
}
