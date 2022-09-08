package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息的内容。
type ConsumeDeadlettersMessageMessage struct {

	// 消息体的内容。
	Body *interface{} `json:"body,omitempty"`

	// 属性的列表。
	Attributes *interface{} `json:"attributes,omitempty"`
}

func (o ConsumeDeadlettersMessageMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeDeadlettersMessageMessage struct{}"
	}

	return strings.Join([]string{"ConsumeDeadlettersMessageMessage", string(data)}, " ")
}
