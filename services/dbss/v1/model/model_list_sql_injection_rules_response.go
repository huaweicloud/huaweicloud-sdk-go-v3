package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlInjectionRulesResponse Response Object
type ListSqlInjectionRulesResponse struct {

	// SQL规则列表
	Rules *[]SqlRuleResponseRules `json:"rules,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlInjectionRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlInjectionRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlInjectionRulesResponse", string(data)}, " ")
}
