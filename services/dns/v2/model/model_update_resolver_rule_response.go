package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResolverRuleResponse Response Object
type UpdateResolverRuleResponse struct {
	ResolverRule   *ResolverRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateResolverRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolverRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateResolverRuleResponse", string(data)}, " ")
}
