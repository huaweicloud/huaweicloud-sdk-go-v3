package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsistencyTaskDetailResponse Response Object
type ShowConsistencyTaskDetailResponse struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业描述
	Description *string `json:"description,omitempty"`

	// 所属目录ID
	CategoryId *int64 `json:"category_id,omitempty"`

	// SUGGEST:提示, MINOR:一般, MAJOR:严重, FATAL:致命
	Level *string `json:"level,omitempty"`

	// 统一告警条件
	AlarmCondition *string `json:"alarm_condition,omitempty"`

	// 是否开启通知告警
	AlarmNotify *bool `json:"alarm_notify,omitempty"`

	// TRIGGER_ALARM:触发告警, RUN_SUCCESS:运行成功, TRIGGER_ALARM_AND_RUNNING_SUCCESS:触发告警和运行成功
	AlarmNotifyType *string `json:"alarm_notify_type,omitempty"`

	// 通知主题名
	AlarmNotifyTopic *string `json:"alarm_notify_topic,omitempty"`

	// 调度类型，ONCE：单次调度，PERIODIC：周期性调度
	ScheduleType *string `json:"schedule_type,omitempty"`

	// 调度周期，MINUTE:按分钟调度，HOUR:按小时调度，DAY:按天调度，WEEK:按周调度
	SchedulePeriod *string `json:"schedule_period,omitempty"`

	// 调度间隔，注意：当调度周期为分钟、小时、天时，间隔时间为数字；而当调度周期为周时，调度间隔为星期的英文，如：每周一、周二调度时，schedule_interval为MONDAY,TUESDAY
	ScheduleInterval *string `json:"schedule_interval,omitempty"`

	// 调度开始时间
	ScheduleStartTime *string `json:"schedule_start_time,omitempty"`

	// 调度结束时间
	ScheduleEndTime *string `json:"schedule_end_time,omitempty"`

	// 最近运行时间,13位时间戳(精确到毫秒)
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最近运行时间,13位时间戳(精确到毫秒)
	LastRunTime *int64 `json:"last_run_time,omitempty"`

	// 子规则
	SubRules       *[][]ConsistencyRuleDetailForOpenApi `json:"sub_rules,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowConsistencyTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsistencyTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowConsistencyTaskDetailResponse", string(data)}, " ")
}
