package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResolveRuleResponse struct {

	// 查询resolver_rule的列表响应。
	ResolverRules  *[]ResolveRuleParam `json:"resolver_rules,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o DeleteResolveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolveRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteResolveRuleResponse", string(data)}, " ")
}
