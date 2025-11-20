package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateForwardRuleResponse Response Object
type UpdateForwardRuleResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateForwardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateForwardRuleResponse", string(data)}, " ")
}
