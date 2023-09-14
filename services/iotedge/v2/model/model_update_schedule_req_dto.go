package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleReqDto 更新调度计划请求结构体
type UpdateScheduleReqDto struct {

	// 调度计划名称
	Name string `json:"name"`

	// 调度计划是否生效
	Enabled bool `json:"enabled"`

	// 调度计划起始时间，毫秒级别的时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 调度计划结束时间，毫秒级别的时间戳
	EndTime int64 `json:"end_time"`

	// 调度计划优先级。
	Priority int32 `json:"priority"`

	Daily *DailyDto `json:"daily,omitempty"`

	// 调度任务信息
	Tasks []ScheduleTask `json:"tasks"`
}

func (o UpdateScheduleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleReqDto struct{}"
	}

	return strings.Join([]string{"UpdateScheduleReqDto", string(data)}, " ")
}
