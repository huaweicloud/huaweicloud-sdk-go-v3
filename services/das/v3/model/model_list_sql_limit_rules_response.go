package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlLimitRulesResponse Response Object
type ListSqlLimitRulesResponse struct {

	// SQL限流规则列表
	SqlLimitRules *[]SqlLimitRule `json:"sql_limit_rules,omitempty"`

	// SQL限流规则总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlLimitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlLimitRulesResponse", string(data)}, " ")
}
