package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclRuleOrderResponse Response Object
type UpdateAclRuleOrderResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAclRuleOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclRuleOrderResponse struct{}"
	}

	return strings.Join([]string{"UpdateAclRuleOrderResponse", string(data)}, " ")
}
