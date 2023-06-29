package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmSubsResponse Response Object
type ListAlarmSubsResponse struct {

	// 告警订阅总数
	Count *int32 `json:"count,omitempty"`

	// 告警订阅列表
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
