package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventScheduleWindow **参数解释**：  执行时间窗。  **约束限制**：  当operation_type为reservation时，该字段必传。
type EventScheduleWindow struct {

	// **参数解释**：  执行日期。  **约束限制**：  格式为“yyyy-mm-dd”。  **取值范围**：  大于或等于当前日期。  **默认取值**：  不涉及。
	PlannedDay string `json:"planned_day"`

	// **参数解释**：  事件执行窗口开始时间。  **约束限制**：  格式为 “hh:mm”。当end_time有值时，该字段必传。  **取值范围**：  不涉及。  **默认取值**：  “01:00”。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  事件执行窗口结束时间。  **约束限制**：  格式为 \"hh:mm\"。当start_time有值时，该字段必传。  **取值范围**：  不涉及。  **默认取值**：  “03:00”。
	EndTime *string `json:"end_time,omitempty"`
}

func (o EventScheduleWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventScheduleWindow struct{}"
	}

	return strings.Join([]string{"EventScheduleWindow", string(data)}, " ")
}
