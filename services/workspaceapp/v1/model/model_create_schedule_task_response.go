package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleTaskResponse Response Object
type CreateScheduleTaskResponse struct {

	// 定时任务主键id。
	Id *string `json:"id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	LastStatus *ScheduleTaskStatus `json:"last_status,omitempty"`

	TaskType *ScheduleTaskTypeEnum `json:"task_type,omitempty"`

	// 定时任务表达式。
	TaskCron *string `json:"task_cron,omitempty"`

	// 下一次执行时间。
	NextExecutionTime *string `json:"next_execution_time,omitempty"`

	ScheduleTaskPolicy *ScheduleTaskPolicy `json:"schedule_task_policy,omitempty"`

	ScheduledType *ScheduledTypeEnum `json:"scheduled_type,omitempty"`

	// 周期按天时：按x天间隔执行。
	DayInterval *int32 `json:"day_interval,omitempty"`

	// 周期按周时：取值1~7，英文逗号分隔，如1,2,7。
	WeekList *string `json:"week_list,omitempty"`

	// 周期按月时：取值1~12，英文逗号分隔，如1,3,12。
	MonthList *string `json:"month_list,omitempty"`

	// 周期按月时：取值1~31及L(代表当月最后一天)，英文逗号分隔，如1,2,28,L。
	DateList *string `json:"date_list,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 周期指定时间时：表示指定的日期，格式为yyyy-MM-dd。
	ScheduledDate *string `json:"scheduled_date,omitempty"`

	// 指定的执行时间点，格式为HH:mm:ss。
	ScheduledTime *string `json:"scheduled_time,omitempty"`

	// 到期时间。
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`

	// 任务描述。
	Description *string `json:"description,omitempty"`

	// 是否是开启状态。
	IsEnable *bool `json:"is_enable,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskResponse", string(data)}, " ")
}
