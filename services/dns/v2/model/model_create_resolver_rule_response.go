package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolverRuleResponse Response Object
type CreateResolverRuleResponse struct {
	ResolverRule   *ResolverRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateResolverRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateResolverRuleResponse", string(data)}, " ")
}
