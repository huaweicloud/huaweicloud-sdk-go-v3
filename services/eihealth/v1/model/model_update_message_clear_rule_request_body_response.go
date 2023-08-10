package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageClearRuleRequestBodyResponse Response Object
type UpdateMessageClearRuleRequestBodyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMessageClearRuleRequestBodyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageClearRuleRequestBodyResponse struct{}"
	}

	return strings.Join([]string{"UpdateMessageClearRuleRequestBodyResponse", string(data)}, " ")
}
