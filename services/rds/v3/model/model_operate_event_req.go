package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateEventReq struct {

	// **参数解释**：  事件列表  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EventInstances []EventInstance `json:"event_instances"`

	// **参数解释**：  事件操作类型：cancel|execute|reservation  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	OperationType string `json:"operation_type"`

	EventScheduleWindow *EventScheduleWindow `json:"event_schedule_window,omitempty"`
}

func (o OperateEventReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateEventReq struct{}"
	}

	return strings.Join([]string{"OperateEventReq", string(data)}, " ")
}
