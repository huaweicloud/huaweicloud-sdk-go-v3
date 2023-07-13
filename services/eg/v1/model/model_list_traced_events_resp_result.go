package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTracedEventsRespResult struct {
	SubscriptionName *string `json:"subscription_name,omitempty"`

	SourceName *string `json:"source_name,omitempty"`

	SourceProvider *string `json:"source_provider,omitempty"`

	EventType *string `json:"event_type,omitempty"`

	EventId *string `json:"event_id,omitempty"`

	EventReceivedTime *int32 `json:"event_received_time,omitempty"`
}

func (o ListTracedEventsRespResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTracedEventsRespResult struct{}"
	}

	return strings.Join([]string{"ListTracedEventsRespResult", string(data)}, " ")
}
