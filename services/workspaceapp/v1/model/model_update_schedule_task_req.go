package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleTaskReq 更新定时任务
type UpdateScheduleTaskReq struct {
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

	// 任务名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成，不能有空格 2. 长度范围1~64个字符
	TaskName *string `json:"task_name,omitempty"`

	TaskType *ScheduleTaskTypeEnum `json:"task_type,omitempty"`

	ScheduleTaskPolicy *ScheduleTaskPolicy `json:"schedule_task_policy,omitempty"`

	// 任务描述。
	Description *string `json:"description,omitempty"`

	// 是否启用： 'true': 启用 'false': 禁用
	IsEnable *bool `json:"is_enable,omitempty"`

	// 定时任务对象列表。
	TargetInfos *[]TargetInfo `json:"target_infos,omitempty"`
}

func (o UpdateScheduleTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleTaskReq struct{}"
	}

	return strings.Join([]string{"UpdateScheduleTaskReq", string(data)}, " ")
}
