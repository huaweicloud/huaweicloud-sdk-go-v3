package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobEventsResponse Response Object
type ListTrainingJobEventsResponse struct {

	// **参数解释**：总条数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：最大显示条数。 **取值范围**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：开始的条数。 **取值范围**：不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：排序方式。 **取值范围**：不涉及。
	Order *string `json:"order,omitempty"`

	// **参数解释**：事件的开始时间。 **取值范围**：不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：事件的结束时间。 **取值范围**：不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：事件列表。
	Events         *[]Event `json:"events,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListTrainingJobEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobEventsResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobEventsResponse", string(data)}, " ")
}
