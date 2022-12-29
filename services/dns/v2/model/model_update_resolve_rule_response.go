package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateResolveRuleResponse struct {
	ResolverRule   *ResolveRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateResolveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolveRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateResolveRuleResponse", string(data)}, " ")
}
