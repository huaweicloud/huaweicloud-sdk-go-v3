package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledEventsResponse Response Object
type ListScheduledEventsResponse struct {

	// **参数解释**：计划事件列表
	Events *[]ScheduledEvent `json:"events,omitempty"`

	// **参数解释**：计划事件总数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListScheduledEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsResponse", string(data)}, " ")
}
