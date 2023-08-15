package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventPubMetricsItem 事件通道监控指标数据
type EventPubMetricsItem struct {

	// 事件大小
	EventSize *int64 `json:"event_size,omitempty"`

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 调用数
	Num *int64 `json:"num,omitempty"`

	// 调用成功数
	SuccessNum *int64 `json:"success_num,omitempty"`

	// 处理时间
	ProcessTime *int64 `json:"process_time,omitempty"`

	// 调用时间
	InvokeTime *int64 `json:"invoke_time,omitempty"`
}

func (o EventPubMetricsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventPubMetricsItem struct{}"
	}

	return strings.Join([]string{"EventPubMetricsItem", string(data)}, " ")
}
