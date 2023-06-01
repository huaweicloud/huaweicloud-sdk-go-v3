package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTracedEventsRequest struct {

	// 查询数据起始时间
	StartTime int64 `json:"start_time"`

	// 查询数据结束时间
	EndTime int64 `json:"end_time"`

	// 指定查询事件的Id
	EventId *string `json:"event_id,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`

	// 指定查询的事件类型，精准匹配
	EventType *string `json:"event_type,omitempty"`

	// 事件订阅名称
	SubscriptionName *string `json:"subscription_name,omitempty"`

	// 每页显示的条目数量，不能小于0
	Limit *string `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 通道ID
	ChannelId string `json:"channel_id"`
}

func (o ListTracedEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTracedEventsRequest struct{}"
	}

	return strings.Join([]string{"ListTracedEventsRequest", string(data)}, " ")
}
