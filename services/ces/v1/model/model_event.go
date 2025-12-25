package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Event 响应参数。
type Event struct {

	// **参数解释**： 事件ID。 **取值范围**： 不涉及。
	EventId string `json:"event_id"`

	// **参数解释**： 事件名称。 **取值范围**： 不涉及。
	EventName string `json:"event_name"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
