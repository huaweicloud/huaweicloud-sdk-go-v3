package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumeMessage struct {
	Message *ConsumeMessageMessage `json:"message,omitempty"`
	// 消息handler。

	Handler *string `json:"handler,omitempty"`
}

func (o ConsumeMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeMessage struct{}"
	}

	return strings.Join([]string{"ConsumeMessage", string(data)}, " ")
}
