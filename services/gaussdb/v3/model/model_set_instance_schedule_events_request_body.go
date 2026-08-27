package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceScheduleEventsRequestBody **参数解释**：  设置事件执行策略参数体。  **约束限制**：  不涉及。
type SetInstanceScheduleEventsRequestBody struct {

	// **参数解释**：  事件操作类型。  **约束限制**：  不涉及。  **取值范围**：  - execute：授权立即执行。 - cancel：授权取消执行。 - reservation：授权预约执行。  **默认取值**：  不涉及。
	OperationType string `json:"operation_type"`

	// **参数解释**：  事件信息。  **约束限制**：  批量事件个数不得超过20个。
	EventInstances []EventInstances `json:"event_instances"`

	// **参数解释**：  事件执行窗口。  **约束限制**：  当operation_type为reservation时，该字段必传。
	EventScheduleWindow *interface{} `json:"event_schedule_window,omitempty"`
}

func (o SetInstanceScheduleEventsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceScheduleEventsRequestBody struct{}"
	}

	return strings.Join([]string{"SetInstanceScheduleEventsRequestBody", string(data)}, " ")
}
