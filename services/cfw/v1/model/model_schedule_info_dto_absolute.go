package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleInfoDtoAbsolute **参数解释**： 绝对时间计划 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type ScheduleInfoDtoAbsolute struct {

	// **参数解释**： 绝对计划开始时间 **约束限制**： 不涉及 **取值范围**： 时间戳（毫秒级） **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 绝对计划结束时间 **约束限制**： 不涉及 **取值范围**： 时间戳（毫秒级） **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ScheduleInfoDtoAbsolute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleInfoDtoAbsolute struct{}"
	}

	return strings.Join([]string{"ScheduleInfoDtoAbsolute", string(data)}, " ")
}
