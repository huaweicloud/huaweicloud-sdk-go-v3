package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleReqDto 创建调度计划请求结构体
type CreateScheduleReqDto struct {

	// 调度计划id，租户下唯一，选填如不填则随机生成
	ScheduleId *string `json:"schedule_id,omitempty"`

	// 调度计划名称
	Name string `json:"name"`

	// 调度计划的循环类型，once表示在start_time执行，end_time结束；daliy表示start_time-end_time之间每天都执行
	CycleType string `json:"cycle_type"`

	// 调度计划是否生效
	Enabled bool `json:"enabled"`

	// 调度计划起始时间，毫秒级别的时间戳，可选值，不填表示立即执行
	StartTime *int64 `json:"start_time,omitempty"`

	// 调度计划结束时间，毫秒级别的时间戳
	EndTime int64 `json:"end_time"`

	// 调度计划优先级。
	Priority int32 `json:"priority"`

	Daily *DailyDto `json:"daily,omitempty"`

	// 调度任务信息
	Tasks []ScheduleTask `json:"tasks"`
}

func (o CreateScheduleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleReqDto struct{}"
	}

	return strings.Join([]string{"CreateScheduleReqDto", string(data)}, " ")
}
