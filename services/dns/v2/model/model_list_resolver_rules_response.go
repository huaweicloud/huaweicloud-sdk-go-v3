package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolverRulesResponse Response Object
type ListResolverRulesResponse struct {

	// 解析器转发规则列表。
	ResolverRules *[]ListResolverRuleParam `json:"resolver_rules,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResolverRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverRulesResponse struct{}"
	}

	return strings.Join([]string{"ListResolverRulesResponse", string(data)}, " ")
}
