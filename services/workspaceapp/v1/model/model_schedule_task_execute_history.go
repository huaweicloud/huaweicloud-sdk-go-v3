package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleTaskExecuteHistory 定时任务执行记录。
type ScheduleTaskExecuteHistory struct {

	// 定时任务执行记录主键id。
	Id *string `json:"id,omitempty"`

	// 定时任务主键id。
	TaskId *string `json:"task_id,omitempty"`

	TaskType *ScheduleTaskTypeEnum `json:"task_type,omitempty"`

	Status *ScheduleTaskStatus `json:"status,omitempty"`

	ScheduledType *ScheduledTypeEnum `json:"scheduled_type,omitempty"`

	// 总子任务数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 成功的子任务数。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败的子任务数。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 正在执行的子任务数。
	ExcutingCount *int32 `json:"excuting_count,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 任务开始时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o ScheduleTaskExecuteHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTaskExecuteHistory struct{}"
	}

	return strings.Join([]string{"ScheduleTaskExecuteHistory", string(data)}, " ")
}
