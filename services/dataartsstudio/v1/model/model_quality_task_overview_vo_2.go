package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityTaskOverviewVo2 struct {

	// id
	Id *int64 `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 目录ID
	CategoryId *int64 `json:"category_id,omitempty"`

	// 调度状态 UNKNOWN:未知,NOT_START:未启动,SCHEDULING:调度中,FINISH_SUCCESS:正常结束,KILL:手动停止,RUNNING_EXCEPTION:运行失败
	ScheduleStatus *string `json:"schedule_status,omitempty"`

	// 调度周期 MINUTE:按分钟调度,HOUR:按小时调度,DAY:按天调度,WEEK:按周调度
	SchedulePeriod *string `json:"schedule_period,omitempty"`

	// 调度间隔 当调度周期为分钟、小时、天时，返回数值字符串，当调度周期为周时，返回具体的调度星期信息如（MONDAY,THURSDAY）
	ScheduleInterval *string `json:"schedule_interval,omitempty"`

	// 创建时间,13位时间戳(精确到毫秒)
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最新运行时间,13位时间戳(精确到毫秒)
	LastRunTime *int64 `json:"last_run_time,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`
}

func (o QualityTaskOverviewVo2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityTaskOverviewVo2 struct{}"
	}

	return strings.Join([]string{"QualityTaskOverviewVo2", string(data)}, " ")
}
