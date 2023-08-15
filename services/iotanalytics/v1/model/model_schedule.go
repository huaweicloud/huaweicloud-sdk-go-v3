package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Schedule SQL作业。
type Schedule struct {

	// 调度开始时间。时间格式为ISO时区日期时间。例如2021-03-03T10:15:30+08:00
	StartTime string `json:"start_time"`

	// 调度结束时间。时间格式为ISO时区日期时间。例如2021-03-03T10:15:30+08:00
	EndTime string `json:"end_time"`

	// Cron表达式，格式为<秒> <分> <时> <天> <月> <星期>
	CronExpression string `json:"cron_expression"`

	// 计算资源ID。
	ComputingResourceId string `json:"computing_resource_id"`

	// 调度启用状态. true: 调度中；false：停止调度。
	Enable bool `json:"enable"`

	// 作业运行配置信息。
	Conf *[]string `json:"conf,omitempty"`

	// 仅在查询作业和查询所有作业接口的响应返回。调度状态。1:NORMAL, 2:PAUSED, 3:COMPLETE, 4:ERROR, 5:BLOCKED
	ScheduleStatus *string `json:"schedule_status,omitempty"`

	// 仅在查询作业和查询所有作业接口的响应返回。上一次调度开始时间。
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// 仅在查询作业和查询所有作业接口的响应返回。下一次调度开始时间。
	PrevFireTime *string `json:"prev_fire_time,omitempty"`
}

func (o Schedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Schedule struct{}"
	}

	return strings.Join([]string{"Schedule", string(data)}, " ")
}
