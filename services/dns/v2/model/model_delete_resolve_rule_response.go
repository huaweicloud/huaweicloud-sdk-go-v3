package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResolveRuleResponse struct {
	ResolverRule   *ResolveRuleParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteResolveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolveRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteResolveRuleResponse", string(data)}, " ")
}
