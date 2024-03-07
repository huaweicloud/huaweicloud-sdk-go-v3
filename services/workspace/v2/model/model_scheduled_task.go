package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTask 定时任务信息。
type ScheduledTask struct {

	// 任务id。
	Id *string `json:"id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型。START：开机，STOP：关机，REBOOT：重启，HIBERNATE：休眠，REBUILD：重建系统盘，EXECUTE_SCRIPT：执行脚本。
	TaskType *string `json:"task_type,omitempty"`

	// 执行周期。FIXED_TIME：指定时间，DAY：按天，WEEK：按周，MONTH：按月。
	ScheduledType *string `json:"scheduled_type,omitempty"`

	// 触发场景类型。
	LifeCycleType *string `json:"life_cycle_type,omitempty"`

	// 最近一次执行状态。SUCCESS：成功，SKIP：跳过，FAIL：失败。
	LastStatus *string `json:"last_status,omitempty"`

	// 下一次执行时间。格式为yyyy-MM-dd HH:mm:ss。
	NextExecutionTime *string `json:"next_execution_time,omitempty"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 优先级。触发式任务使用。
	Priority *int32 `json:"priority,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ScheduledTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTask struct{}"
	}

	return strings.Join([]string{"ScheduledTask", string(data)}, " ")
}
