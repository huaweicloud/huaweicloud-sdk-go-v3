package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfirmConsumptionMessagesReq struct {
	// 确认消息数组。

	Message *[]ConfirmMessageEntity `json:"message,omitempty"`
}

func (o ConfirmConsumptionMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmConsumptionMessagesReq struct{}"
	}

	return strings.Join([]string{"ConfirmConsumptionMessagesReq", string(data)}, " ")
}
