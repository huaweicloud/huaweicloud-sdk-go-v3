package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsResponse Response Object
type ListEventsResponse struct {

	// **参数解释**： 事件详情列表。 **取值范围**： 不涉及。
	Events *[]EventResponse `json:"events,omitempty"`

	// **参数解释**： 事件总数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
