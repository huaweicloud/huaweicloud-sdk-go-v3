package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchExecuteEventsRequestBody struct {

	// **参数解释**: 事件列表。 **约束限制**: 不涉及。
	EventInstances []EventInstanceOption `json:"event_instances"`

	// **参数解释**: 事件操作类型。 **约束限制**: 不涉及。 **取值范围**: - cancel：取消事件 - execute：立即执行 - reservation：预约执行时间窗口  **默认取值**: 不涉及。
	OperationType string `json:"operation_type"`

	EventScheduleWindow *EventScheduleWindowOption `json:"event_schedule_window,omitempty"`
}

func (o BatchExecuteEventsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteEventsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchExecuteEventsRequestBody", string(data)}, " ")
}
