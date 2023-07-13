package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageClearRuleRequest Request Object
type UpdateMessageClearRuleRequest struct {
	Body *SetMessageClearRuleReq `json:"body,omitempty"`
}

func (o UpdateMessageClearRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageClearRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageClearRuleRequest", string(data)}, " ")
}
