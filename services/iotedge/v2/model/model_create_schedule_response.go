package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleResponse Response Object
type CreateScheduleResponse struct {

	// 调度计划id，租户下唯一
	ScheduleId *string `json:"schedule_id,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 调度计划名称
	Name *string `json:"name,omitempty"`

	// 调度计划的循环类型
	CycleType *string `json:"cycle_type,omitempty"`

	// 调度计划是否生效
	Enabled *bool `json:"enabled,omitempty"`

	// 调度计划起始时间，毫秒级别的时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 调度计划结束时间，毫秒级别的时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 调度计划优先级, 1-16
	Priority *int32 `json:"priority,omitempty"`

	Daily *DailyDto `json:"daily,omitempty"`

	// 调度任务信息
	Tasks          *[]ScheduleTask `json:"tasks,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduleResponse", string(data)}, " ")
}
