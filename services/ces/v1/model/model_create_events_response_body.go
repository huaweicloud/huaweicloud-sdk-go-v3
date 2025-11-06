package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventsResponseBody 响应参数。
type CreateEventsResponseBody struct {

	// **参数解释**： 事件ID。 **取值范围**： 不涉及。
	EventId string `json:"event_id"`

	// **参数解释**： 事件名称。 **取值范围**： 不涉及。
	EventName string `json:"event_name"`
}

func (o CreateEventsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsResponseBody struct{}"
	}

	return strings.Join([]string{"CreateEventsResponseBody", string(data)}, " ")
}
