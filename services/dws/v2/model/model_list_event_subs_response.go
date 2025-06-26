package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSubsResponse Response Object
type ListEventSubsResponse struct {

	// **参数解释**： 事件订阅总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 事件订阅详情列表。 **取值范围**： 不涉及。
	EventSubscriptions *[]EventSubscriptionResponse `json:"event_subscriptions,omitempty"`
	HttpStatusCode     int                          `json:"-"`
}

func (o ListEventSubsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSubsResponse struct{}"
	}

	return strings.Join([]string{"ListEventSubsResponse", string(data)}, " ")
}
