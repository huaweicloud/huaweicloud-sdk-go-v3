package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventScheduleWindow struct {

	// **参数解释**：  计划执行日期 yyyy-MM-dd  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	PlannedDay string `json:"planned_day"`

	// **参数解释**：  计划执行窗口开始时间  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  计划执行窗口结束时间  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o EventScheduleWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventScheduleWindow struct{}"
	}

	return strings.Join([]string{"EventScheduleWindow", string(data)}, " ")
}
