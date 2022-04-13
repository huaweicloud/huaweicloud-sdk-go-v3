package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息的内容。
type ConsumeMessageMessage struct {
	// 消息体的内容。

	Body *interface{} `json:"body,omitempty"`
	// 属性的列表。

	Attributes *interface{} `json:"attributes,omitempty"`
	// 标签值。

	Tags *[]string `json:"tags,omitempty"`
}

func (o ConsumeMessageMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeMessageMessage struct{}"
	}

	return strings.Join([]string{"ConsumeMessageMessage", string(data)}, " ")
}
