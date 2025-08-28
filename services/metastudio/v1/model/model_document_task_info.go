package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentTaskInfo 文档分段任务信息。
type DocumentTaskInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务进度
	TaskProcess *string `json:"task_process,omitempty"`

	// 任务结果
	TaskResult *string `json:"task_result,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 任务创建时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	// 任务更新时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o DocumentTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentTaskInfo struct{}"
	}

	return strings.Join([]string{"DocumentTaskInfo", string(data)}, " ")
}
