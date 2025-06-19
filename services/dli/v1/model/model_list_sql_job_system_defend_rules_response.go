package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobSystemDefendRulesResponse Response Object
type ListSqlJobSystemDefendRulesResponse struct {

	// 系统预制规则信息。
	Rules *[]ShowSqlJobSystemDefendRuleResponseBody `json:"rules,omitempty"`

	// 总个数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlJobSystemDefendRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobSystemDefendRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlJobSystemDefendRulesResponse", string(data)}, " ")
}
