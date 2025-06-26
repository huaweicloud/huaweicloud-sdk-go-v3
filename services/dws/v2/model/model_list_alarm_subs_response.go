package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmSubsResponse Response Object
type ListAlarmSubsResponse struct {

	// **参数解释**： 告警订阅总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 告警订阅列表。 **取值范围**： 不涉及。
	AlarmSubscriptions *[]AlarmSubscriptionResponse `json:"alarm_subscriptions,omitempty"`
	HttpStatusCode     int                          `json:"-"`
}

func (o ListAlarmSubsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmSubsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmSubsResponse", string(data)}, " ")
}
