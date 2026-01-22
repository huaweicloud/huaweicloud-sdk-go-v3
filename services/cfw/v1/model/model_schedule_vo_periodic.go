package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleVoPeriodic struct {

	// **参数解释**： 周期计划类型 **取值范围**： 0：每天 1：每周的某几天 2：每周的一段时间
	Type *int32 `json:"type,omitempty"`

	// **参数解释**： 周期计划开始时间 **取值范围**： 不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 周期计划结束时间 **取值范围**： 不涉及
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 每周的某几天 **取值范围**： 不涉及
	WeekMask *[]int32 `json:"week_mask,omitempty"`

	// **参数解释**： 每周周期计划开始日 **取值范围**： 不涉及
	StartWeek *int32 `json:"start_week,omitempty"`

	// **参数解释**： 每周周期计划结束日 **取值范围**： 不涉及
	EndWeek *int32 `json:"end_week,omitempty"`
}

func (o ScheduleVoPeriodic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleVoPeriodic struct{}"
	}

	return strings.Join([]string{"ScheduleVoPeriodic", string(data)}, " ")
}
