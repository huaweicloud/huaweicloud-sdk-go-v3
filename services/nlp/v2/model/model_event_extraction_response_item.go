package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// events数据结构
type EventExtractionResponseItem struct {
	// 事件元素列表。

	Argument []EventExtractionResponseItemArgument `json:"argument"`
	// 事件触发词。触发词是事件描述中最能代表事件发生的词汇，决定事件类别的重要特征。

	EventTrigger string `json:"event_trigger"`
	// 事件类型。

	EventType *string `json:"event_type,omitempty"`
	// 事件触发词在待分析文本中的起始和终止位置。

	TriggerSpan []int32 `json:"trigger_span"`
}

func (o EventExtractionResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventExtractionResponseItem struct{}"
	}

	return strings.Join([]string{"EventExtractionResponseItem", string(data)}, " ")
}
