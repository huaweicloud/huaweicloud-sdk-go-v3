package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTasksRecords 定时任务执行记录。
type ScheduledTasksRecords struct {

	// 任务执行记录id
	Id *string `json:"id,omitempty"`

	// 执行时间，格式为yyyy-MM-dd HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 任务类型。START：开机，STOP：关机，REBOOT：重启，HIBERNATE：休眠，REBUILD：重建系统盘。
	TaskType *string `json:"task_type,omitempty"`

	// 执行周期类型。FIXED_TIME：指定时间，DAY：按天，WEEK：按周，MONTH：按月。
	ScheduledType *string `json:"scheduled_type,omitempty"`

	// 本次执行状态。
	Status *string `json:"status,omitempty"`

	// 成功桌面个数。
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败桌面个数。
	FailedNum *int32 `json:"failed_num,omitempty"`

	// 跳过桌面个数。
	SkipNum *int32 `json:"skip_num,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 执行定时任务的任务id，只有定时执行脚本返回。
	ExecuteTaskId *string `json:"execute_task_id,omitempty"`
}

func (o ScheduledTasksRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTasksRecords struct{}"
	}

	return strings.Join([]string{"ScheduledTasksRecords", string(data)}, " ")
}
