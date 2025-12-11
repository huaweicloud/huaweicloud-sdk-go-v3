package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlRuleRankResponse Response Object
type SetSqlRuleRankResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSqlRuleRankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlRuleRankResponse struct{}"
	}

	return strings.Join([]string{"SetSqlRuleRankResponse", string(data)}, " ")
}
