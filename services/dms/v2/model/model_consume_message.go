package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumeMessage struct {
	Message *ConsumeMessageMessage `json:"message,omitempty" xml:"message"`

	// 消息handler。
	Handler *string `json:"handler,omitempty" xml:"handler"`
}

func (o ConsumeMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeMessage struct{}"
	}

	return strings.Join([]string{"ConsumeMessage", string(data)}, " ")
}
