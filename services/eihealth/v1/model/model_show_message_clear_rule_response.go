package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageClearRuleResponse Response Object
type ShowMessageClearRuleResponse struct {

	// 最多保留记录数
	MessageRetainNumber *int32 `json:"message_retain_number,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowMessageClearRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageClearRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowMessageClearRuleResponse", string(data)}, " ")
}
