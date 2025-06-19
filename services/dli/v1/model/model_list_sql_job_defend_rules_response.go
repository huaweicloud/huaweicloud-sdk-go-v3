package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobDefendRulesResponse Response Object
type ListSqlJobDefendRulesResponse struct {

	// 用户自定义规则信息。
	Rules *[]SqlJobDefendRule `json:"rules,omitempty"`

	// 总个数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlJobDefendRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobDefendRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlJobDefendRulesResponse", string(data)}, " ")
}
