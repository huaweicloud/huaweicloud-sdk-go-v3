package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResolverRuleResponse Response Object
type DeleteResolverRuleResponse struct {
	ResolverRule   *ResolverRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o DeleteResolverRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolverRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteResolverRuleResponse", string(data)}, " ")
}
