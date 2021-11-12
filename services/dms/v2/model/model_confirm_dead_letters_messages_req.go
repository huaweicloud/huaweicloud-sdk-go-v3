package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfirmDeadLettersMessagesReq struct {
	// 确认消息数组。

	Message *[]ConfirmDeadLettersMessagesReqMessage `json:"message,omitempty"`
}

func (o ConfirmDeadLettersMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDeadLettersMessagesReq struct{}"
	}

	return strings.Join([]string{"ConfirmDeadLettersMessagesReq", string(data)}, " ")
}
