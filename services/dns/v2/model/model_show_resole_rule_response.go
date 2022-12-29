package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResoleRuleResponse struct {
	ResolverRule   *ShowResolveRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowResoleRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResoleRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowResoleRuleResponse", string(data)}, " ")
}
