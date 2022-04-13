package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfirmDeadLettersMessagesReq struct {
	// 确认消息数组。

	Message *[]ConfirmMessageEntity `json:"message,omitempty"`
}

func (o ConfirmDeadLettersMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDeadLettersMessagesReq struct{}"
	}

	return strings.Join([]string{"ConfirmDeadLettersMessagesReq", string(data)}, " ")
}
