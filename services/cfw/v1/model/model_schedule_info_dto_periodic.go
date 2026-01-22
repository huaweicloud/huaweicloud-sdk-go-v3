package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleInfoDtoPeriodic struct {

	// **参数解释**： 周期计划类型 **约束限制**： 不涉及 **取值范围**： 0：每天 1：每周的某几天 2：每周的一段时间 **默认取值**： 不涉及
	Type int32 `json:"type"`

	// **参数解释**： 周期计划开始时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	StartTime string `json:"start_time"`

	// **参数解释**： 周期计划结束时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EndTime string `json:"end_time"`

	// **参数解释**： 每周的某几天 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	WeekMask *[]int32 `json:"week_mask,omitempty"`

	// **参数解释**： 每周周期计划开始日 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	StartWeek *int32 `json:"start_week,omitempty"`

	// **参数解释**： 每周周期计划结束日 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EndWeek *int32 `json:"end_week,omitempty"`
}

func (o ScheduleInfoDtoPeriodic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleInfoDtoPeriodic struct{}"
	}

	return strings.Join([]string{"ScheduleInfoDtoPeriodic", string(data)}, " ")
}
