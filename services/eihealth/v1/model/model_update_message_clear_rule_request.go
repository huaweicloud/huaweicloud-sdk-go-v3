package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMessageClearRuleRequest struct {
	Body *GetMessageClearRuleReq `json:"body,omitempty"`
}

func (o UpdateMessageClearRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageClearRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageClearRuleRequest", string(data)}, " ")
}
