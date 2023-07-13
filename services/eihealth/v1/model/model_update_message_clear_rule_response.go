package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageClearRuleResponse Response Object
type UpdateMessageClearRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMessageClearRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageClearRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateMessageClearRuleResponse", string(data)}, " ")
}
