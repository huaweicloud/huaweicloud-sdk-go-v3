package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateResolveRuleResponse struct {

	// 查询resolver_rule的列表响应。
	ResolverRules  *[]ResolveRuleParam `json:"resolver_rules,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateResolveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolveRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateResolveRuleResponse", string(data)}, " ")
}
