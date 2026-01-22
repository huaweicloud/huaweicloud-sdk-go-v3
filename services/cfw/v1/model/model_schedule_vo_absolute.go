package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleVoAbsolute **参数解释**： 绝对时间计划 **取值范围**： 不涉及
type ScheduleVoAbsolute struct {

	// **参数解释**： 绝对计划开始时间 **取值范围**： 不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 绝对计划结束时间 **取值范围**： 不涉及
	EndTime *string `json:"end_time,omitempty"`
}

func (o ScheduleVoAbsolute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleVoAbsolute struct{}"
	}

	return strings.Join([]string{"ScheduleVoAbsolute", string(data)}, " ")
}
