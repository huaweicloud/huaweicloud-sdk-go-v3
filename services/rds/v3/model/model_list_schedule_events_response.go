package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleEventsResponse Response Object
type ListScheduleEventsResponse struct {

	// **参数解释**：  事件总数。  **约束限制**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  待授权的事件数。  **约束限制**：  不涉及。
	InquiringCount *int32 `json:"inquiring_count,omitempty"`

	// **参数解释**：  待执行的事件数。  **约束限制**：  不涉及。
	ScheduleCount *int32 `json:"schedule_count,omitempty"`

	// **参数解释**：  正在执行的事件数。  **约束限制**：  不涉及。
	ExecutingCount *int32 `json:"executing_count,omitempty"`

	// **参数解释**：  执行失败的事件数。  **约束限制**：  不涉及。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// **参数解释**：  事件详情列表。  **约束限制**：  不涉及。
	Events         *[]ScheduleEventInfo `json:"events,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListScheduleEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleEventsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleEventsResponse", string(data)}, " ")
}
