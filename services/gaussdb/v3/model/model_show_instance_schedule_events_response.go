package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceScheduleEventsResponse Response Object
type ShowInstanceScheduleEventsResponse struct {

	// **参数解释**：  事件总数。  **取值范围**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  待授权的事件数。  **取值范围**：  不涉及。
	InquiringCount *int32 `json:"inquiring_count,omitempty"`

	// **参数解释**：  待执行的事件数。  **取值范围**：  不涉及。
	ScheduleCount *int32 `json:"schedule_count,omitempty"`

	// **参数解释**：  正在执行的事件数。  **取值范围**：  不涉及。
	ExecutingCount *int32 `json:"executing_count,omitempty"`

	// **参数解释**：  执行失败的事件数。  **取值范围**：  不涉及。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// **参数解释**：  事件详情列表，包含事件ID、事件类别、事件状态、事件级别、实例信息、执行时间等详细信息
	Events         *[]ScheduleEventInfo `json:"events,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowInstanceScheduleEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceScheduleEventsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceScheduleEventsResponse", string(data)}, " ")
}
