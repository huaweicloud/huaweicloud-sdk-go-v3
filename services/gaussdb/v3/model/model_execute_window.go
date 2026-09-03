package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteWindow **参数解释**：  事件执行窗口。  **约束限制**：  不涉及。
type ExecuteWindow struct {

	// **参数解释**：  计划执行日期，格式为：“yyyy-MM-dd”。  **取值范围**：  不涉及。
	PlannedExecutionDay *string `json:"planned_execution_day,omitempty"`

	// **参数解释**：  事件执行窗口开始时间，格式为 \"hh:mm\"。  **取值范围**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  事件执行窗口结束时间，格式为 \"hh:mm\"。  **取值范围**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ExecuteWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteWindow struct{}"
	}

	return strings.Join([]string{"ExecuteWindow", string(data)}, " ")
}
