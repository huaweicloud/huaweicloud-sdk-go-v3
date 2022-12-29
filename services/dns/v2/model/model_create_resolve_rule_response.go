package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateResolveRuleResponse struct {
	ResolverRule   *ResolveRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateResolveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolveRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateResolveRuleResponse", string(data)}, " ")
}
