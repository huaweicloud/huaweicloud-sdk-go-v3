package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAclRuleResponse Response Object
type AddAclRuleResponse struct {
	Data           *RuleIdList `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o AddAclRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAclRuleResponse struct{}"
	}

	return strings.Join([]string{"AddAclRuleResponse", string(data)}, " ")
}
