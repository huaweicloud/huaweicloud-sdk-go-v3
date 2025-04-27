package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResolverRuleResponse Response Object
type ShowResolverRuleResponse struct {
	ResolverRule   *ShowResolverRuleRespParam `json:"resolver_rule,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowResolverRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResolverRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowResolverRuleResponse", string(data)}, " ")
}
