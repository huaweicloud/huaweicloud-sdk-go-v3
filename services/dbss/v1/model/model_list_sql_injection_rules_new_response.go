package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlInjectionRulesNewResponse Response Object
type ListSqlInjectionRulesNewResponse struct {

	// SQL规则列表
	Rules *[]SqlRuleResponseRules `json:"rules,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlInjectionRulesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlInjectionRulesNewResponse struct{}"
	}

	return strings.Join([]string{"ListSqlInjectionRulesNewResponse", string(data)}, " ")
}
