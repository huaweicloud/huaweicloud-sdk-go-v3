package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResoleRulesResponse struct {

	// 查询resolver_rule的列表响应。
	ResolverRules  *[]ResolveRuleParam `json:"resolver_rules,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListResoleRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResoleRulesResponse struct{}"
	}

	return strings.Join([]string{"ListResoleRulesResponse", string(data)}, " ")
}
