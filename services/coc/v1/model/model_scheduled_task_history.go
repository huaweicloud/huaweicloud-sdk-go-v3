package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTaskHistory 定时运维历史记录
type ScheduledTaskHistory struct {

	// 历史记录ID
	Id *string `json:"id,omitempty"`

	// 引用任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 执行ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 引用任务名称
	AssociatedTaskName *string `json:"associated_task_name,omitempty"`

	// 引用任务名称(英文)
	AssociatedTaskNameEn *string `json:"associated_task_name_en,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 开始时间时间戳
	StartedTime *int64 `json:"started_time,omitempty"`

	// 结束时间时间戳
	FinishedTime *int64 `json:"finished_time,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 执行结果描述
	ExecutionMsg *string `json:"execution_msg,omitempty"`
}

func (o ScheduledTaskHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskHistory struct{}"
	}

	return strings.Join([]string{"ScheduledTaskHistory", string(data)}, " ")
}
