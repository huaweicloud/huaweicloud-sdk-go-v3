package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleConfig **参数解释：** 创建服务请求体。 **取值范围：** 不涉及。
type ScheduleConfig struct {

	// **参数解释：** 对应的时间单位的数值。 **约束限制：** 与time_unit共同确认时间设置的范围是1分钟~7天之间。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释：** 时间的单位。 **约束限制：** 与duration共同确认时间设置的范围是1分钟~7天之间。 **取值范围：** - MINUTES：分钟。 - HOURS：小时。 - DAYS：天。 **默认取值：** 不涉及。
	TimeUnit *string `json:"time_unit,omitempty"`

	// **参数解释：** 调度类型，当前仅支持取值为STOP。 **约束限制：** 不涉及。 **取值范围：** - STOP：停止。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ScheduleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleConfig struct{}"
	}

	return strings.Join([]string{"ScheduleConfig", string(data)}, " ")
}
