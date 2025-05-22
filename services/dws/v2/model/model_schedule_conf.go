package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleConf 调度配置信息
type ScheduleConf struct {

	// **参数解释**： 调度开始时间。 **取值范围**： 不涉及。
	ScheduleStart *string `json:"schedule_start,omitempty"`

	// **参数解释**： 调度结束时间。 **取值范围**： 不涉及。
	ScheduleEnd *string `json:"schedule_end,omitempty"`

	// **参数解释**： 调度类型。 **取值范围**： 不涉及。
	ScheduleType *string `json:"schedule_type,omitempty"`

	// **参数解释**： 调度日期。 **取值范围**： 不涉及。
	ScheduleDate *[]int32 `json:"schedule_date,omitempty"`

	// **参数解释**： 调度时间列表。 **取值范围**： 不涉及。
	ScheduleTime *[]string `json:"schedule_time,omitempty"`
}

func (o ScheduleConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleConf struct{}"
	}

	return strings.Join([]string{"ScheduleConf", string(data)}, " ")
}
