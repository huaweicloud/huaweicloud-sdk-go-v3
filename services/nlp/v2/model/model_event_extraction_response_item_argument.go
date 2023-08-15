package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventExtractionResponseItemArgument 事件元素
type EventExtractionResponseItemArgument struct {

	// 元素角色。元素角色指的是事件元素在事件中扮演的角色，是事件元素与事件的语义关系。
	Role string `json:"role"`

	// 实体文本在待分析文本中的起始和终止位置。
	Span []int32 `json:"span"`

	// 实体文本。
	Word string `json:"word"`
}

func (o EventExtractionResponseItemArgument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventExtractionResponseItemArgument struct{}"
	}

	return strings.Join([]string{"EventExtractionResponseItemArgument", string(data)}, " ")
}
