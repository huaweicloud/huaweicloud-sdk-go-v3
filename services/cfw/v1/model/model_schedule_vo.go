package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleVo struct {

	// **参数解释**： 时间表ID **取值范围**： 不涉及
	ScheduleId *string `json:"schedule_id,omitempty"`

	// **参数解释**： 时间表名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 时间表描述 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 引用次数 **取值范围**： 不涉及
	RefCount *int32 `json:"ref_count,omitempty"`

	// **参数解释**： 周期计划 **取值范围**： 不涉及
	Periodic *[]ScheduleVoPeriodic `json:"periodic,omitempty"`

	Absolute *ScheduleVoAbsolute `json:"absolute,omitempty"`
}

func (o ScheduleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleVo struct{}"
	}

	return strings.Join([]string{"ScheduleVo", string(data)}, " ")
}
