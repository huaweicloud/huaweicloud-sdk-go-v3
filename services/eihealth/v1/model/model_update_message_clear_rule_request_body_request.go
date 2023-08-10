package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageClearRuleRequestBodyRequest Request Object
type UpdateMessageClearRuleRequestBodyRequest struct {
	Body *SetMessageClearRuleReq `json:"body,omitempty"`
}

func (o UpdateMessageClearRuleRequestBodyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageClearRuleRequestBodyRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageClearRuleRequestBodyRequest", string(data)}, " ")
}
