package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclRuleResponse Response Object
type DeleteAclRuleResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAclRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAclRuleResponse", string(data)}, " ")
}
