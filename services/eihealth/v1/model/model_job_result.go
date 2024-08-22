package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResult 作业运行结果信息
type JobResult struct {

	// 输入总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 失败个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 子任务运行时长（秒）。
	SubTasksDuration *[]float32 `json:"sub_tasks_duration,omitempty"`
}

func (o JobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResult struct{}"
	}

	return strings.Join([]string{"JobResult", string(data)}, " ")
}
