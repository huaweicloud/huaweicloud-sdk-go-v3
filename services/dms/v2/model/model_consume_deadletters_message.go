package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumeDeadlettersMessage struct {
	Message *ConsumeDeadlettersMessageMessage `json:"message,omitempty"`
	// 消息handler。

	Handler *string `json:"handler,omitempty"`
}

func (o ConsumeDeadlettersMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeDeadlettersMessage struct{}"
	}

	return strings.Join([]string{"ConsumeDeadlettersMessage", string(data)}, " ")
}
