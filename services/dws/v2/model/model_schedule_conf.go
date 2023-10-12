package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleConf 调度配置信息
type ScheduleConf struct {

	// 调度开始时间
	ScheduleStart *string `json:"schedule_start,omitempty"`

	// 调度结束时间
	ScheduleEnd *string `json:"schedule_end,omitempty"`

	// 调度类型
	ScheduleType *string `json:"schedule_type,omitempty"`

	// 调度日期
	ScheduleDate *[]int32 `json:"schedule_date,omitempty"`

	// 调度时间列表
	ScheduleTime *[]string `json:"schedule_time,omitempty"`
}

func (o ScheduleConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleConf struct{}"
	}

	return strings.Join([]string{"ScheduleConf", string(data)}, " ")
}
