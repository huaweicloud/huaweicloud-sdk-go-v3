package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleConfigResponse **参数解释：** 定时停止配置。
type ScheduleConfigResponse struct {

	// **参数解释：** 触发时间，UTC毫秒，13位时间戳。 **取值范围：** 不涉及。
	DueTime *int64 `json:"due_time,omitempty"`

	// **参数解释：** 对应的时间单位的数值。 **取值范围：** 不涉及。
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释：** 调度时间单位。 **取值范围：** - MINUTES：分钟。 - HOURS：小时。 - DAYS：天。
	TimeUnit *string `json:"time_unit,omitempty"`

	// **参数解释：** 调度类型，当前仅支持取值为STOP。 **取值范围：** - STOP：停止。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 表示是否处理完成。 **取值范围：** - true：该定时任务已经执行过。 - false：该定时任务尚未执行。
	Processed *bool `json:"processed,omitempty"`
}

func (o ScheduleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleConfigResponse struct{}"
	}

	return strings.Join([]string{"ScheduleConfigResponse", string(data)}, " ")
}
