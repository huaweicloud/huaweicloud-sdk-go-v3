package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleResponse Response Object
type UpdateRuleResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleResponse", string(data)}, " ")
}
