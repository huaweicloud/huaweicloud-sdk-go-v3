package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventSubsResponse struct {

	// 事件订阅总数
	Count *int32 `json:"count,omitempty"`

	// 事件订阅详情列表
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
